import gzip
import requests
import uuid
from fastapi import FastAPI
from pydantic import BaseModel
class Item(BaseModel):
    url: str
    data:str
app = FastAPI()
@app.post("/upload")
async def upload_request(item: Item):
    headers = {"ProtoType": "json", "brkey": str(uuid.uuid1()), "Br-Content-Encoding": "gzip"}
    upload_response = requests.post(url=item.url, headers=headers, data=gzip.compress(bytes(item.data,"utf-8")))
    return gzip.decompress(upload_response.content).decode()
