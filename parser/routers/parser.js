import { Router } from 'express'
import * as controllers from '../controllers/parser.js'

export const router = Router()

const use = fn => async (req,res,next) => {
    Promise.resolve(fn(req,res,next)).catch(next)
}

router.get('/exchangers', use(controllers.ParseData))
router.get('/exchangers/:city', use(controllers.ParseDataByCity))

// router.get('/exchangers/:city', use(controllers.ParseExchangersByCity))
// router.get('/exchangers', use(controllers.ParseAllExchangers))

// router.get('/currencies/:city', use(controllers.ParseCurrenciesByCity))
// router.get('/currencies', use(controllers.ParseAllCurrencies))
