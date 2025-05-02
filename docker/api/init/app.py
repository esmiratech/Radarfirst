import grpc
import helloworld_pb2 as hello_pb2
import helloworld_pb2_grpc as hello_pb2_grpc
import os
from flask import Flask, jsonify, request
from flask_cors import CORS, cross_origin
import logging
import time

app = Flask(__name__)
cors = CORS(app)
app.config['CORS_HEADERS'] = 'Content-Type'

# Configure logging
logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)

grpc_endpoint = os.environ["GRPC_ENDPOINT"]

@app.before_request
def before_request():
    request.start_time = time.time()
    logger.info(f"request started: {request.method} {request.path}")


@app.after_request
def after_request(response):
    end_time = time.time()
    elapsed_time = end_time - request.start_time
    logger.info(f"request completed: {request.method} {request.path}. Elapsed time: {elapsed_time:.2f}s")
    return response

@app.route('/hello', methods=['POST'])
@cross_origin()
def say_hello():
    data = request.get_json()
    name = data.get('name', '')
    if not name:
        return jsonify({'error': 'Name is required'}), 400
    message = send_gRPC_request(name)
    return jsonify({'message': message})


def send_gRPC_request(name):
    channel = grpc.insecure_channel(grpc_endpoint)
    stub = hello_pb2_grpc.GreeterStub(channel)
    request = hello_pb2.HelloRequest(name=name)
    response = stub.SayHello(request)
    return response.message

if __name__ == '__main__':
    app.run(port=8080, host='0.0.0.0')
