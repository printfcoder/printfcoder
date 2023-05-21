-- 比尔盖茨基金
-- 代码、名称、持有人、周期、数量
SELECT sb.dm, sb.mc, st10g.holder_name, st10g.hold_amount, st10g.end_date
FROM stock_base sb
         LEFT JOIN stock_top_10_gudong st10g ON st10g.dm = sb.dm
WHERE st10g.holder_name LIKE 'BILL %'
  AND end_date = '20230331'
ORDER BY end_date;


-- 查询每天没有同步完成的daily
SELECT count(1)
FROM stock_base sb
         LEFT JOIN
     stock_qt_daily sqt ON sqt.dai_ma = sb.dm
         AND sqt.shi_jian > 20230519000001 AND sqt.shi_jian < 20230519235959
WHERE sqt.dai_ma IS NULL
  -- and sqt2.dang_qian_jia_ge <>0;

