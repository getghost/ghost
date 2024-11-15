import { ClickHouse } from "@ghost/clickhouse";
import { env } from "./env";

export const clickhouse = new ClickHouse({ url: env().CLICKHOUSE_URL });
