import axios from "axios";
import puppeteer from "puppeteer";
import { LAUNCH_PUPPETEER_OPTS, PAGE_PUPPETEER_OPTS } from './config.js';

export const pageContent = async (url) => {
    const browser = await puppeteer.launch(LAUNCH_PUPPETEER_OPTS);
    const page = await browser.newPage();

    await page.goto(url);
    const content = await page.content();
    await browser.close()

    return content
}