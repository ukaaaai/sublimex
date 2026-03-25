import torch
from sentence_transformers import SentenceTransformer
from pb import embedding_pb2, embedding_pb2_grpc

class EmbeddingService(embedding_pb2_grpc.EmbeddingServiceServicer):
    def __init__(self):
        device = "cuda" if torch.cuda.is_available() else "cpu"
        self.model = SentenceTransformer("cl-nagoya/ruri-v3-310m", device=device)
    
    def GetEmbeddings(self, request : embedding_pb2.EmbeddingRequest, context) -> embedding_pb2.EmbeddingResponse:
        embeddings = self.model.encode(request.texts)

        response = embedding_pb2.EmbeddingResponse()
        for emb in embeddings:
            vector = embedding_pb2.Vector(values=emb.tolist())
            response.vectors.append(vector)
            
        return response
