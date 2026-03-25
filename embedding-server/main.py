import grpc
from concurrent import futures
from pb import embedding_pb2_grpc
from service import embedding_service

def main():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    embedding_pb2_grpc.add_EmbeddingServiceServicer_to_server(embedding_service.EmbeddingService)
    server.add_insecure_port('[::]:50051')
    print("🚀 gRPC server started on port 50051")
    server.start()
    server.wait_for_termination()

if __name__ == "__main__":
    main()
