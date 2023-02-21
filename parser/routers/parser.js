import { Router } from 'express'
import * as controllers from '../controllers/parser.js'

export const router = Router()

const use = fn => async (req,res,next) => {
    Promise.resolve(fn(req,res,next)).catch(next)
}

router.get('/exchangers', use(controllers.ParseData))
router.get('/exchangers/:city', use(controllers.ParseDataByCity))
