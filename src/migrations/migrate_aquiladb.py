# this script is used to migrate data from old aquilaDB databases to new ones 
# when a new ML encoder model is added in aquila hub.

from cassandra.cluster import Cluster

import base58, base64, uuid, time
import requests
import json

from aquilapy import Wallet, DB

# Create a wallet instance from private key
wallet = Wallet("/ossl/private_unencrypted.pem")

# Connect to Aquila DB instance
db = DB("http://aquiladb", "5001", wallet)

def create_database (user_id):

    # Schema definition to be used
    schema_def = {
        "description": "Wikipedia",
        "unique": user_id,
        "encoder": "ftxt:https://ftxt-models.s3.us-east-2.amazonaws.com/cc.en.300.bin",
        "codelen": 768,
        "metadata": {
            "url": "string",
            "text": "string"
        }
    }

    # Craete a database with the schema definition provided
    db_name = db.create_database(schema_def)

    return db_name, True

# create a cassandra reader session
def create_session (clusters_arr, kspace):
    cluster = Cluster(clusters_arr)
    # try connecting
    try:
        session = cluster.connect()
        session.default_timeout = 120
    except Exception as e:
        print(e)
        return None
    # try setting keyspace
    try:
        session.set_keyspace(kspace)
    except Exception as e:
        print(e)
        # create session keyspace
        session.execute("CREATE KEYSPACE "+kspace+" \
            WITH replication = {'class':'SimpleStrategy', 'replication_factor' : 1};")
        # set keyspace
        session.set_keyspace(kspace)

    return session

def create_temp_dbs (logging_session, user_session):
    query1 = "CREATE TABLE IF NOT EXISTS content_index_by_database_t ( \
            id_ varint, \
            database_name varchar, \
            url text, \
            html text, \
            timestamp varint, \
            is_deleted int, \
            PRIMARY KEY ((database_name), timestamp, id_) ) \
            WITH CLUSTERING ORDER BY ( timestamp DESC, id_ ASC );"

    query2 = "CREATE TABLE IF NOT EXISTS content_metadata_by_database_t ( \
            id_ varint, \
            database_name varchar, \
            url text, \
            coverimg text, \
            title text, \
            author text, \
            timestamp varint, \
            outlinks text, \
            summary text, \
            PRIMARY KEY ((database_name), timestamp, id_) ) \
            WITH CLUSTERING ORDER BY ( timestamp DESC, id_ ASC );"

    query3 = "CREATE TABLE IF NOT EXISTS search_history_by_database_t ( \
            id_ varint, \
            database_name varchar, \
            query text, \
            url text, \
            timestamp varint, \
            PRIMARY KEY ((database_name), timestamp, id_) ) \
            WITH CLUSTERING ORDER BY ( timestamp DESC, id_ ASC );"

    query4 = "CREATE TABLE IF NOT EXISTS search_correction_by_database_t ( \
            id_ varint, \
            database_name varchar, \
            query text, \
            url text, \
            timestamp varint, \
            PRIMARY KEY ((database_name), timestamp, id_) ) \
            WITH CLUSTERING ORDER BY ( timestamp DESC, id_ ASC );"

    query5 = "CREATE TABLE IF NOT EXISTS search_index_by_user_t ( \
            usecret varchar, \
            aquila_database_name varchar, \
            pub_db_id varchar, \
            pub_enabled int, \
            is_deleted int, \
            timestamp varint, \
            PRIMARY KEY ((usecret), timestamp, aquila_database_name) ) \
            WITH CLUSTERING ORDER BY ( timestamp DESC, aquila_database_name ASC );"

    query6 = "CREATE TABLE IF NOT EXISTS user_profile_by_email_t ( \
            usecret varchar, \
            email varchar, \
            name varchar, \
            title text, \
            avatar_url text, \
            is_deleted int, \
            timestamp varint, \
            PRIMARY KEY ((email), timestamp, usecret) ) \
            WITH CLUSTERING ORDER BY ( timestamp DESC, usecret ASC );"

    query7 = "CREATE TABLE IF NOT EXISTS public_subscribe_list_by_user_t ( \
            usecret varchar, \
            is_deleted int, \
            timestamp varint, \
            pub_db_id varchar, \
            PRIMARY KEY ((usecret), timestamp, pub_db_id) ) \
            WITH CLUSTERING ORDER BY ( timestamp DESC, pub_db_id ASC );"
            
    try:
        logging_session.execute(query1)
        logging_session.execute(query2)
        logging_session.execute(query3)
        logging_session.execute(query4)
        user_session.execute(query5)
        user_session.execute(query6)
        user_session.execute(query7)

        return True
    except Exception as e:
        print(e)
        return False


def wipe_old_temp_dbs (logging_session, user_session):
    query1 = "TRUNCATE content_index_by_database_t;"

    query2 = "TRUNCATE content_metadata_by_database_t;"

    query3 = "TRUNCATE search_history_by_database_t;"

    query4 = "TRUNCATE search_correction_by_database_t;"

    query5 = "TRUNCATE search_index_by_user_t;"

    query6 = "TRUNCATE user_profile_by_email_t;"

    query7 = "TRUNCATE public_subscribe_list_by_user_t;"
            
    try:
        logging_session.execute(query1)
        logging_session.execute(query2)
        logging_session.execute(query3)
        logging_session.execute(query4)
        user_session.execute(query5)
        user_session.execute(query6)
        user_session.execute(query7)

        return True
    except Exception as e:
        print(e)
        return False

def copy_to_temp_dbs (logging_session, user_session):
    try:
        # direct copy contents
        res = user_session.execute("SELECT * FROM user_profile_by_email ALLOW FILTERING;")
        for r in res:
            user_session.execute("INSERT INTO user_profile_by_email_t (usecret, email, name, title, avatar_url, is_deleted, timestamp) \
                VALUES('{}', '{}', '{}', '{}', '{}', {}, {});".format(r.usecret, r.email, r.name, r.title, r.avatar_url, r.is_deleted, r.timestamp))
        
        res = user_session.execute("SELECT * FROM public_subscribe_list_by_user ALLOW FILTERING;")
        for r in res:
            user_session.execute("INSERT INTO public_subscribe_list_by_user_t (usecret, pub_db_id, is_deleted, timestamp) \
            VALUES('{}', '{}', {}, {});".format(r.usecret, r.pub_db_id, r.is_deleted, r.timestamp))

        # create new db names for each users
        adb_old_new_map = {}
        res = user_session.execute("SELECT * FROM search_index_by_user ALLOW FILTERING;")
        for r in res:
            if not adb_old_new_map.get(r.aquila_database_name):
                seed = base58.b58encode(uuid.uuid4().bytes)[:-14].decode("utf-8")+str(int(time.time()))
                db_name, status = create_database(seed)
                if not status:
                    return False
                adb_old_new_map[r.aquila_database_name] = db_name
        
        res = user_session.execute("SELECT * FROM search_index_by_user ALLOW FILTERING;")
        for r in res:
            user_session.execute("INSERT INTO search_index_by_user_t (usecret, aquila_database_name, pub_db_id, pub_enabled, is_deleted, timestamp) \
            VALUES('{}', '{}', '{}', {}, {}, {});".format(r.usecret, adb_old_new_map[r.aquila_database_name], r.pub_db_id, r.pub_enabled, r.is_deleted, r.timestamp))
        
        res = logging_session.execute("SELECT * FROM content_index_by_database ALLOW FILTERING;")
        for r in res:
            logging_session.execute("INSERT INTO content_index_by_database_t (id_, database_name, url, html, timestamp, is_deleted) \
            VALUES({}, '{}', '{}', '{}', {}, {});".format(r.id_, adb_old_new_map[r.database_name], r.url, r.html, r.timestamp, r.is_deleted))
        
        res = logging_session.execute("SELECT * FROM content_metadata_by_database ALLOW FILTERING;")
        for r in res:
            logging_session.execute("INSERT INTO content_metadata_by_database_t (id_, database_name, url, coverimg, title, author, timestamp, outlinks, summary) \
            VALUES({}, '{}', '{}', '{}', '{}', '{}', {}, '{}', '{}');".format(r.id_, adb_old_new_map[r.database_name], r.url, r.coverimg, r.title, r.author, r.timestamp, r.outlinks, r.summary))
        
        res = logging_session.execute("SELECT * FROM search_history_by_database ALLOW FILTERING;")
        for r in res:
            logging_session.execute("INSERT INTO search_history_by_database_t (id_, database_name, query, url, timestamp) \
            VALUES({}, '{}', '{}', '{}', {});".format(r.id_, adb_old_new_map[r.database_name], r.query, r.url, r.timestamp))
        
        res = logging_session.execute("SELECT * FROM search_correction_by_database ALLOW FILTERING;")
        for r in res:
            logging_session.execute("INSERT INTO search_correction_by_database_t (id_, database_name, query, url, timestamp) \
            VALUES({}, '{}', '{}', '{}', {});".format(r.id_, adb_old_new_map[r.database_name], r.query, r.url, r.timestamp))
        
        return True
    except Exception as e:
        print(e)
        return False


def wipe_old_dbs (logging_session, user_session):
    query1 = "TRUNCATE content_index_by_database;"

    query2 = "TRUNCATE content_metadata_by_database;"

    query3 = "TRUNCATE search_history_by_database;"

    query4 = "TRUNCATE search_correction_by_database;"

    query5 = "TRUNCATE search_index_by_user;"

    query6 = "TRUNCATE user_profile_by_email;"

    query7 = "TRUNCATE public_subscribe_list_by_user;"
            
    try:
        logging_session.execute(query1)
        logging_session.execute(query2)
        logging_session.execute(query3)
        logging_session.execute(query4)
        user_session.execute(query5)
        user_session.execute(query6)
        user_session.execute(query7)

        return True
    except Exception as e:
        print(e)
        return False

def copy_back_from_temp_dbs (logging_session, user_session):
    try:
        # direct copy contents
        res = user_session.execute("SELECT * FROM user_profile_by_email_t ALLOW FILTERING;")
        for r in res:
            user_session.execute("INSERT INTO user_profile_by_email (usecret, email, name, title, avatar_url, is_deleted, timestamp) \
                VALUES('{}', '{}', '{}', '{}', '{}', {}, {});".format(r.usecret, r.email, r.name, r.title, r.avatar_url, r.is_deleted, r.timestamp))
        
        res = user_session.execute("SELECT * FROM public_subscribe_list_by_user_t ALLOW FILTERING;")
        for r in res:
            user_session.execute("INSERT INTO public_subscribe_list_by_user (usecret, pub_db_id, is_deleted, timestamp) \
            VALUES('{}', '{}', {}, {});".format(r.usecret, r.pub_db_id, r.is_deleted, r.timestamp))
        
        res = user_session.execute("SELECT * FROM search_index_by_user_t ALLOW FILTERING;")
        for r in res:
            user_session.execute("INSERT INTO search_index_by_user (usecret, aquila_database_name, pub_db_id, pub_enabled, is_deleted, timestamp) \
            VALUES('{}', '{}', '{}', {}, {}, {});".format(r.usecret, r.aquila_database_name, r.pub_db_id, r.pub_enabled, r.is_deleted, r.timestamp))
        
        res = logging_session.execute("SELECT * FROM content_index_by_database_t ALLOW FILTERING;")
        for r in res:
            logging_session.execute("INSERT INTO content_index_by_database (id_, database_name, url, html, timestamp, is_deleted) \
            VALUES({}, '{}', '{}', '{}', {}, {});".format(r.id_, r.database_name, r.url, r.html, r.timestamp, r.is_deleted))
        
        res = logging_session.execute("SELECT * FROM content_metadata_by_database_t ALLOW FILTERING;")
        for r in res:
            logging_session.execute("INSERT INTO content_metadata_by_database (id_, database_name, url, coverimg, title, author, timestamp, outlinks, summary) \
            VALUES({}, '{}', '{}', '{}', '{}', '{}', {}, '{}', '{}');".format(r.id_, r.database_name, r.url, r.coverimg, r.title, r.author, r.timestamp, r.outlinks, r.summary))
        
        res = logging_session.execute("SELECT * FROM search_history_by_database_t ALLOW FILTERING;")
        for r in res:
            logging_session.execute("INSERT INTO search_history_by_database (id_, database_name, query, url, timestamp) \
            VALUES({}, '{}', '{}', '{}', {});".format(r.id_, r.database_name, r.query, r.url, r.timestamp))
        
        res = logging_session.execute("SELECT * FROM search_correction_by_database_t ALLOW FILTERING;")
        for r in res:
            logging_session.execute("INSERT INTO search_correction_by_database (id_, database_name, query, url, timestamp) \
            VALUES({}, '{}', '{}', '{}', {});".format(r.id_, r.database_name, r.query, r.url, r.timestamp))
        
        return True
    except Exception as e:
        print(e)
        return False

def populate_aquilaDB (logging_session):
    url = "http://localhost:5003/index"
    headers = {
        'Content-Type': 'application/json'
    }
    try:
        res = logging_session.execute("SELECT * FROM content_index_by_database_t ALLOW FILTERING;")
        for r in res:
            print("---")
            # r.database_name, r.url, r.html
            payload = json.dumps({
                "database": r.database_name,
                "html": base64.b64decode(r.html.encode("utf-8")).decode("utf-8"),
                "url": r.url
            })
            response = requests.request("POST", url, headers=headers, data=payload)
            print(response)

        return True
    except Exception as e:
        print(e)
        return False

if __name__ == "__main__":
    logging_session = create_session(["164.52.214.80"], 'logging')
    user_session = create_session(["164.52.214.80"], 'users')
    time.sleep(1)
    print(create_temp_dbs(logging_session, user_session))
    time.sleep(1)
    print(wipe_old_temp_dbs(logging_session, user_session))
    time.sleep(1)
    print(copy_to_temp_dbs(logging_session, user_session))
    time.sleep(1)
    print(populate_aquilaDB(logging_session))
    # time.sleep(1)
    # print(wipe_old_dbs(logging_session, user_session))
    # time.sleep(1)
    # print(copy_back_from_temp_dbs(logging_session, user_session))
