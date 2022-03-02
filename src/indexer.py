import logging

import html_cleanup as chtml

import metadata_parser

import threading
import queue
import time

class Indexer ():
    def __init__(self, db, hub):
        self.db = db
        self.hub = hub
        self.q_maxsize = 100
        self.process_flag = True
        self.process_timeout_sec = 1 # in seconds

        # spawn process thread
        self.spawn()

    def __del__(self):
        self.process_flag = False
        if self.process_thread:
            self.process_thread.join()

    # index webpage parent
    def index (self, html_data, url, db_name):
        # fetch metadata
        page_metadata = metadata_parser.MetadataParser(html=html_data, search_head_only=True)
        # fill in metadata object to reflect mercury's response
        chtml_data = {"data": {}}
        # title
        if page_metadata.get_metadatas('title') != None:
            chtml_data["data"]["title"] = page_metadata.get_metadatas('title')[0]
        else:
            chtml_data["data"]["title"] = url
        # author
        if page_metadata.get_metadatas('author') != None:
            chtml_data["data"]["author"] = page_metadata.get_metadatas('author')[0]
        else:
            chtml_data["data"]["author"] = None
        # cover image
        if page_metadata.get_metadatas('image') != None:
            chtml_data["data"]["lead_image_url"] = page_metadata.get_metadatas('image')[0]
        else:
            chtml_data["data"]["lead_image_url"] = None
        # summary
        if page_metadata.get_metadatas('description') != None:
            chtml_data["data"]["excerpt"] = page_metadata.get_metadatas('description')[0]
        else:
            chtml_data["data"]["excerpt"] = ""
        # next page
        chtml_data["data"]["next_page_url"] = None

        # add to index html queue
        self.pipeline.put({"db_name":db_name, "html_data": html_data, "url":url})

        return True, chtml_data

    
    # Insert website as Aquila DB docs
    def index_website (self, db_name, paragraphs, title, url):
        # add title as well
        if title != "":
            paragraphs.append(title)
        compressed = self.hub.compress_documents(db_name, paragraphs)
        docs = []
        for idx_, para in enumerate(paragraphs):
            v = compressed[idx_]

            docs.append({
                "metadata": {
                    "url": url, 
                    "text": para
                },
                "code": v
            })
        try:
            dids = self.db.insert_documents(db_name, docs)
            return True
        except Exception as e:
            logging.debug(e)
            return False

    def spawn (self):
        # create pipeline to add documents
        self.pipeline = queue.Queue(maxsize=self.q_maxsize)
        # create process thread
        self.process_thread = threading.Thread(target=self.process, args=(), daemon=True)
        # start process thread
        self.process_thread.start()
        # return self.pipeline

    def process(self):
        while (self.process_flag):
            # set a timeout till next vector indexing
            time.sleep(self.process_timeout_sec)

            # check if queue is not empty
            if self.pipeline.qsize() > 0:
                qitem = self.pipeline.get()

                # cleanup html
                chtml_data = chtml.process_html(qitem["html_data"], qitem["url"])
                thtml_data = chtml.trim_content(chtml_data["data"]["content"])["result"]
                
                status = self.index_website (qitem["db_name"], thtml_data, chtml_data["data"]["title"], qitem["url"])
                if not status:
                    logging.error("website indexing in the backround failed for " + qitem["url"])
