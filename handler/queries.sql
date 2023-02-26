SELECT
    upload_time AS time,
    avg(usd_buy) as usd_buy,
    avg(usd_sell) as usd_sell
FROM exchangers_currencies
    WHERE usd_buy > 0 AND upload_time >= 1677386645
GROUP BY upload_time
    ORDER BY upload_time ASC;

