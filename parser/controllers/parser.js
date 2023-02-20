import * as parser from "../services/parser.js"
import {AllExchangers, ParseExchangersByCity} from "../services/parser.js";

export const ParseData = async(req,res,next) => {
    const data = await parser.GetData()
    res.status(200).json(data)
}

export const ParseDataByCity = async(req,res,next) => {
    const city = req.params.city
    const data = await parser.GetDataByCity()
    res.status(200).json(data)
}


export const ParseAllExchangers = async (req,res,next) => {
    const data = await parser.AllExchangers()
    res.status(200).json(data)
}

export const ParseExchangersByCity = async (req,res,next) => {
    const city = req.params.city
    const data = await parser.ExchangersByCity(city)

    res.status(200).json(data)
}

export const ParseCurrenciesByCity = async (req,res,next) => {
    const city = req.params.city
    const data = await parser.CurrenciesByCity(city)
    res.status(200).json(data)
}

export const ParseAllCurrencies = async (req,res,next) => {
    const data = await parser.AllCurrencies()
    res.status(200).json(data)
}