import axios from "axios";
import fs from 'fs';

import {load} from "cheerio";
import * as data from "./data.js"
import puppeteer from "puppeteer";
import * as configs from '../configs/config.js';

export const pageContent = async (url) => {
    // const browser = await puppeteer.launch(configs.LAUNCH_PUPPETEER_OPTS);
    // const page = await browser.newPage(configs.PAGE_PUPPETEER_OPTS);

    // await page.goto(url);
    // const content = await page.content();
    // await browser.close()

    const content = await axios.get(url, {
            responseType: 'document'
        }).then(response => {
            return response.data;
            // Use the DOM API to parse the XML content
        }).catch(error => {
            console.error(error);
        });


    const $ = load(content)

    const script_text = $('script').text()
    const start_idx = script_text.indexOf("punkts")
    const end_idx = script_text.indexOf("var globalTown")

    if (script_text === 0) {
        return
    }

    let file_text = "export const " + script_text.substring(start_idx, end_idx)
    fs.writeFile('./services/data.js', file_text, (error) => {
        if (error !== null) console.log(error);
    });

    return data.punkts
}

const ParsedWithPup = async () => {
    return
    // const data = []
    console.log($('script').text());
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

}