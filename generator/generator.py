#!/usr/bin/env python3

from datetime import datetime
from http.server import HTTPServer, BaseHTTPRequestHandler
import logging
import os
import uuid

logging.basicConfig(level=logging.INFO)

kube_host = os.getenv("KUBERNETES_SERVICE_HOST")
if kube_host is not None:
    print(f"☸ This binary is probably running on Kubernetes ({kube_host})! WE HATE KUBERNETES! ☸")


class UUIDRequestHandler(BaseHTTPRequestHandler):
    def do_GET(self):
        self.protocol_version = 'HTTP/1.1'
        self.send_response(200)
        self.end_headers()

        id = uuid.uuid4()
        logging.info(f"[{datetime.now()}]: generated {id}")
        self.wfile.write(bytes(str(id), "utf-8"))

logging.info("Listening on 0.0.0.0:3000")
HTTPServer(("0.0.0.0", 3000), UUIDRequestHandler).serve_forever()
