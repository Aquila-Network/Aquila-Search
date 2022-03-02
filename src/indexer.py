import logging

import html_cleanup as chtml

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
        # cleanup html
        chtml_data = chtml.process_html(html_data, url)
        thtml_data = chtml.trim_content(chtml_data["data"]["content"])["result"]

        # add to index html queue
        self.pipeline.put({"db_name":db_name, "paragraphs": thtml_data, "title": chtml_data["data"]["title"], "url":url})

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
                status = self.index_website (qitem["db_name"], qitem["paragraphs"], qitem["title"], qitem["url"])
                if not status:
                    logging.error("website indexing in the backround failed for " + qitem["url"])
