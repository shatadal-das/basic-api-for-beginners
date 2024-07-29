from fastapi import FastAPI

app = FastAPI()

@app.post("/hello")
def handler():
    return "Hello World!"


if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="0.0.0.0", port=5000)