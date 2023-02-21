import * as parser from "../services/parser.js"

export const ParseData = async(req,res,next) => {
    const data = await parser.GetExchangers()
    res.status(200).json(data)
}

export const ParseDataByCity = async(req,res,next) => {
    const city = req.params.city
    console.log(city);
    const data = await parser.GetExchangersByCity(city)
    res.status(200).json(data)
}
