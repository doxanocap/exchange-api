import express, { json } from 'express';

import cors from "cors";
import cookieParser from 'cookie-parser';
import { config } from 'dotenv';
import * as services from './services/exchanger.js'
config()


import { router } from './routers/data.js';
import {InitExchangers} from "./services/exchanger.js";

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

app.use((err, req, res, next) => {
  const status = err.status
  if (status !== undefined) {
    delete err.status
    res.status(status).json(err )
    return
  }
  console.log(err);
  res.status(500).json({error:"Something went wrong"})
})

const port = process.env.apiPort || 8002

const server = app.listen(port, async () => {
  await services.InitExchangers()

  console.log(`Example app listening on port ${port}`)
})

export default server
