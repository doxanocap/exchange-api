import axios from "axios";
import puppeteer from "puppeteer";
import * as configs from '../configs/config.js';

export const pageContent = async (url) => {
    const browser = await puppeteer.launch(configs.LAUNCH_PUPPETEER_OPTS);
    const page = await browser.newPage(configs.PAGE_PUPPETEER_OPTS);

    await page.goto(url);
    const content = await page.content();
    await browser.close()

    return content
}