import logging

from flask import Flask, request
from flask_cors import CORS
from functools import wraps
import hub_proxy as hub

from services import logging as slog
slogging_session = slog.create_session(["172.16.126.158"])

import math
import numpy as np

from aquilapy import Wallet, DB

from transformers import pipeline
summarizer = None # pipeline("summarization")
qa = None # pipeline("question-answering")
def init_augmentation ():
    global summarizer
    global qa
    summarizer = pipeline("summarization")
    qa = pipeline("question-answering")
    return True


app = Flask(__name__, instance_relative_config=True)

# Create a wallet instance from private key
wallet = Wallet("/ossl/private_unencrypted.pem")

# Connect to Aquila DB instance
db = DB("http://aquiladb", "5001", wallet)

# Connect to Aquila Hub instance
# hub = Hub("http://aquilahub", "5002", wallet)

# Create indexer instance with routine thread
from indexer import Indexer
web_indexer = Indexer(db, hub)

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

    # Craete a database with the schema definition provided
    # db_name_ = hub.create_database(schema_def)

    return db_name, True

# generate summary
def summarize(text):
    if not summarizer:
        init_augmentation()
    return summarizer(text[:1024], min_length=5, max_length=20)[0]["summary_text"]

# generate QA
def QAgen(query, context):
    if not qa:
        init_augmentation()
    return qa(question=query, context=context)

# Search docs
def search_docs(db_name, query):
    compressed = hub.compress_documents(db_name, [query])
    docs, dists = db.search_k_documents(db_name, compressed, 100)
    index = {}
    score = {}
    # max_score = dists[0][-1]
    # min_score = dists[0][0]

    for idx_, doc in enumerate(docs[0]):
        metadata = doc["metadata"]
        # -------------------------- exponential dampening ------------------------------
        # ------------------- normalize --------------------------
        #      ------ reposition --------
        #                                                            - 1->0 : lesser steep curve -
        # (1 - (dists[0][idx_]-min_score) / (max_score-min_score)) * math.exp(-0.06*idx_)
        if index.get(metadata["url"]):
            # index[metadata["url"]] += 1
            # score[metadata["url"]] += dists[0][idx_] * math.exp(-0.06*idx_)
            pass
        else:
            index[metadata["url"]] = 1
            score[metadata["url"]] = dists[0][idx_] # * math.exp(-0.06*idx_)

    results_d = {}
    # n_unique_urls = len(index.keys())
    
    for key in index:
        #                           ---- Representative rebalance ----
        results_d[key] = score[key] # / index[key] / n_unique_urls

    results_d = {k: v for k, v in sorted(results_d.items(), key=lambda item: item[1], reverse=True)}

    return results_d

# Add authentication
def authenticate ():
    def decorator (f):
        @wraps(f)
        def wrapper (*args, **kwargs):
            # skip
            return f(*args, **kwargs)

        return wrapper
    return decorator

def extract_request_params (request):
    if not request.is_json:
        logging.error("Cannot parse request parameters")

        # request is invalid
        return {}

    # Extract JSON data
    data_ = request.get_json()

    return data_

@app.route("/", methods=['GET'])
def info ():
    """
    Check server status
    """

    # Build response
    return {
            "success": True,
            "message": "Aquila X is running healthy"
        }, 200

@app.route("/create", methods=['POST'])
@authenticate()
def create_db ():
    """
    Create a database on demand given a random unique seed
    """

    # get parameters
    user_id = None
    if extract_request_params(request).get("seed"):
        user_id = extract_request_params(request)["seed"]

    if not user_id:
        # Build error response
        return {
                "success": False,
                "message": "Invalid parameters"
            }, 400

    db_name, status = create_database(user_id)

    # Build response
    if status:
        return {
                "success": True,
                "databaseName": db_name
            }, 200
    else:
        return {
                "success": False,
                "message": "Invalid schema definition"
            }, 400

@app.route("/index", methods=['POST'])
@authenticate()
def index_page ():
    """
    Index html page
    """

    # get parameters
    html_data = None
    url = None
    db_name = None
    if extract_request_params(request).get("database") and extract_request_params(request).get("html") and extract_request_params(request).get("url"):
        html_data = extract_request_params(request)["html"]
        url = extract_request_params(request)["url"]
        db_name = extract_request_params(request)["database"]

    if not html_data or not url or not db_name:
        # Build error response
        return {
                "success": False,
                "message": "Invalid parameters"
            }, 400

    # index a page
    logging.debug("adding website to queue")
    status, chtml_data = web_indexer.index(html_data, url, db_name)
    logging.debug("added website to queue")

    # Build response
    if status:
        # logging
        if slogging_session != None:
            # index activity logging
            slog.put_log_index(slogging_session, db_name, url, html_data, 0)
            # metadata logging
            slog.put_url_summary(slogging_session, db_name, url, chtml_data["data"]["title"], chtml_data["data"]["author"], chtml_data["data"]["lead_image_url"], chtml_data["data"]["next_page_url"], chtml_data["data"]["excerpt"])
            logging.debug("indexing done")
        return {
                "success": True,
                "databaseName": db_name
            }, 200
    else:
        return {
                "success": False,
                "message": "Invalid schema definition"
            }, 400

@app.route("/search", methods=['POST'])
def search ():
    """
    Search database for matches
    """

    # get parameters
    query = None
    db_name = None
    if extract_request_params(request).get("database") and extract_request_params(request).get("query"):
        db_name = extract_request_params(request)["database"]
        query = extract_request_params(request)["query"]

    if not query or not db_name:
        # Build error response
        return {
                "success": False,
                "message": "Invalid parameters"
            }, 400

    urls = search_docs(db_name, query)

    # logging
    if slogging_session != None:
        if len(urls) > 0:
            slog.put_log_search(slogging_session, db_name, query, list(urls.keys())[0])
        else:
            slog.put_log_search(slogging_session, db_name, query, "")

    # Build response
    return {
            "success": True,
            "result": urls
        }, 200

@app.route("/augment", methods=['POST'])
def augment ():
    """
    Augment matches
    """

    # get parameters
    query = None
    context = None
    if extract_request_params(request).get("query") and extract_request_params(request).get("context"):
        query = extract_request_params(request)["query"]
        context = extract_request_params(request)["context"]

    if not query and not context:
        # Build error response
        return {
                "success": False,
                "message": "Invalid parameters"
            }, 400

    summary_r = summarize(context)
    ans_r = QAgen(query, context)

    # Build response
    return {
            "success": True,
            "result": {
                "summary": summary_r,
                "ans": ans_r
            }
        }, 200

@app.route("/correct", methods=['POST'])
def correct ():
    """
    Correct matches
    """

    # get parameters
    query = None
    db_name = None
    url = None
    if extract_request_params(request).get("query") and extract_request_params(request).get("database") and extract_request_params(request).get("url"):
        query = extract_request_params(request)["query"]
        db_name = extract_request_params(request)["database"]
        url = extract_request_params(request)["url"]

    if not query and not db_name and not url:
        # Build error response
        return {
                "success": False,
                "message": "Invalid parameters"
            }, 400

    # logging
    if slogging_session != None:
        slog.put_log_correct(slogging_session, db_name, query, url)

    # index correction
    status = web_indexer.index_website(db_name, [], query, url)

    # Build response
    return {
            "success": True
        }, 200

@app.route("/list", methods=['POST'])
def listall ():
    """
    List indexed urls
    """

    # get parameters
    page = None
    db_name = None
    limit = None
    if extract_request_params(request).get("page") and extract_request_params(request).get("database") and extract_request_params(request).get("limit"):
        page = extract_request_params(request)["page"]
        db_name = extract_request_params(request)["database"]
        limit = extract_request_params(request)["limit"]

    if not page and not db_name and not limit:
        # Build error response
        return {
                "success": False,
                "message": "Invalid parameters"
            }, 400

    # get links
    if slogging_session != None:
        links = slog.get_all_url(slogging_session, db_name, page, limit)

    # Build response
    return {
            "success": True,
            "result": {
                "links": links
            }
        }, 200



@app.route("/urlsummary", methods=['POST'])
def summary ():
    """
    URL summary
    """

    # get parameters
    urls = None
    if extract_request_params(request).get("urls"):
        urls = extract_request_params(request)["urls"]
        db_name = extract_request_params(request)["database"]

    if not urls:
        # Build error response
        return {
                "success": False,
                "message": "Invalid parameters"
            }, 400

    summary_r = slog.get_url_summary(slogging_session, db_name, urls)

    # Build response
    return {
            "success": True,
            "result": {
                "summary": summary_r
            }
        }, 200


# Server starter
def flaskserver ():
    """
    start server
    """
    app.run(host='0.0.0.0', port=5003, debug=False)

# Enable CORS
CORS(app)

if __name__ == "__main__":
    # create default database
    db_name, status = create_database("default")
    logging.debug("Default DB name: " + db_name)
    if status:
        flaskserver()