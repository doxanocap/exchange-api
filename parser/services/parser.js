import * as handler from "./handle.js";
import {load} from "cheerio";

const cities = ["astana","almaty","taldykorgan", "kostanai"]

export const GetDataByCity = async (city) => {
    const content = await handler.pageContent(`https://kurs.kz/index.php?mode=${city}`)
    const $ = load(content)

    const data = []

    $('.punkt-open').each((i, el) => {

        const exchanger = {
            name:  $(el).find($('.tab')).text(),
            link: $(el).find($('.tab')).attr('href'),
            address: $(el).find($('.address')).text(),
            wholesale: $(el).find($('.wholesale')).text(),
            timeC: $(el).find($('.timeC')).text(),
        }

        const currencies = []
        const phones = []


        $(el).find($('.currency')).each((j, ch) => {
            for (const elem of ch.children) {
                const title = elem.attribs.title
                if (title !== undefined) {
                    currencies.push({
                        title: elem.attribs.title,
                        price: $(elem).text()
                    })
                }
            }
        })

        $(el).find('.phone').each((j, ch) => {
            phones.push($(ch).text())
        });

        exchanger.currencies = currencies
        exchanger.phones = phones

        data.push(exchanger)
    })
    return data
}

export const ExchangersByCity = async (city) => {
    const content = await handler.pageContent(`https://kurs.kz/index.php?mode=${city}`)
    const $ = load(content)

    const exchangers = []

    $('.punkt-open').each((i, el) => {

        const exchanger = {
            name:  $(el).find($('.tab')).text(),
            link: $(el).find($('.tab')).attr('href'),
            address: $(el).find($('.address')).text(),
            wholesale: $(el).find($('.wholesale')).text(),
            timeC: $(el).find($('.timeC')).text(),
        }

        const currencies = []
        const phones = []

        $(el).find('.phone').each((j, ch) => {
            phones.push($(ch).text())
        });

        exchanger.currencies = currencies
        exchanger.phones = phones

        exchangers.push(exchanger)
    })
    return exchangers
}

export const CurrenciesByCity = async (city) => {
    const content = await handler.pageContent(`https://kurs.kz/index.php?mode=${city}`)
    const $ = load(content)
    const currencies = []
    
    $('.punkt-open').each((i, el) => {

        $(el).find($('.currency')).each((j, ch) => {
            for (const elem of ch.children) {
                const title = elem.attribs.title
                if (title !== undefined) {
                    currencies.push({
                        title: elem.attribs.title,
                        price: $(elem).text()
                    })
                }
            }
        })
    })
    return currencies
};

export const AllCurrencies = async () => {
    const data = []
    cities.map(async (item,key) =>  {
        const res = await CurrenciesByExchanger(item)
        data.push(res)
    })
    return data
}

export const AllExchangers = async () => {
    const data = []
    cities.map(async (item,key) =>  {
        const res = await ExchangersByCity(item)
        data.push(res)
    })
    return data
}


export const GetData = async () => {
    const data = []
    cities.map(async (item,key) =>  {
        const res = await DataByCity(item)
        data.push(res)
    })
    return data
}