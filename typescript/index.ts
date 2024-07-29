import express, { Request, Response } from "express";

const app = express();

app.post("/hello", handler);

function handler(req: Request, res: Response) {
  res.send("Hello World!");
}

app.listen(5000, () => {
  console.log("Server is running on port 5000");
});
