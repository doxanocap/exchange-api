import express, { json } from 'express';
import cors from "cors";
import cookieParser from 'cookie-parser';
import { config } from 'dotenv';
import { router } from './routers/parser.js';
import { errorHandler } from './middlewares/error_handler.js';

config()

const corsOptions = {
  exposedHeaders: '*',
  origin: 'http://localhost:3000',
  methods: 'GET, PUT, POST',
  credentials: true,
}

const app = express()

app.use(cookieParser());
app.use(json());
app.use(cors(corsOptions))
app.use('/kzt-parser',router)
app.use(errorHandler)


const port = process.env.apiPort || 8050
const server = app.listen(port, async () => {
  console.log(`Example app listening on port ${port}`)
})

export default server
