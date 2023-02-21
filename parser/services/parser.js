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
    const content = await handler.pageContent(`https://kurs.kz/index.php?mode=${city}`)
    const $ = load(content)

    const data = []

    $('.punkt-open').each((i, el) => {

        const exchanger = {
            city: city,
            name:  $(el).find($('.tab')).text(),
            link: $(el).find($('.tab')).attr('href'),
            address: $(el).find($('.address')).text(),
            special_offer: $(el).find($('.wholesale')).text(),
            update_time: $(el).find($('.timeC')).text(),
        }

        const phones = []
        $(el).find($('.currency')).each((j, ch) => {
            for (const elem of ch.children) {
                const title = elem.attribs.title
                if (title !== undefined) {
                    let price = $(elem).text()

                    switch (elem.attribs.title) {
                        case "USD - покупка":
                            exchanger.USD_BUY = parseFloat(price)
                        case "USD - продажа": 
                            exchanger.USD_SELL = parseFloat(price)
                        case "EUR - покупка": 
                            exchanger.EUR_BUY = parseFloat(price)    
                        case "EUR - продажа": 
                            exchanger.EUR_SELL = parseFloat(price)
                        case "RUB - покупка": 
                            exchanger.RUB_BUY = parseFloat(price)
                        case "RUB - продажа": 
                            exchanger.RUB_SELL = parseFloat(price)
                    }
                   
                }
            }
        })

        $(el).find('.phone').each((j, ch) => {
            phones.push($(ch).text())
        });

        exchanger.phone_numbers = phones

        data.push(exchanger)
    })
    return data
}
