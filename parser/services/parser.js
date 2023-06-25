import * as handler from "./handle.js";
import {load} from "cheerio";


const cities = ["astana","almaty","taldykorgan", "kostanai"]

export const GetExchangers = async () => {
    const data = []
    cities.map(async (item,key) =>  {
        const res = await GetExchangersByCity(item)
        data.push(res)
    })
    return data
}

export const GetExchangersByCity = async (city) => {
    const parsed_data = await handler.pageContent(`https://kurs.kz/index.php?mode=${city}`)
    if (!parsed_data) {
        return
    }
    const response = []
    for (let punkt of parsed_data) {
        let exchanger = {
            id: 0,
            name: punkt.name,
            city: punkt.city,
            address: punkt.address,
            wholesale: punkt.phone,
            updated_time: punkt.date,
            phone_numbers: punkt.phones,
            USD: punkt.data.USD,
            EUR: punkt.data.EUR,
            RUB: punkt.data.RUB,
        }
        response.push(exchanger)
    }
    return response
}
