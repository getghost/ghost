import { ClickHouse } from "@ghost/clickhouse";
import { z } from "zod";

export class Analytics {
  private readonly clickhouse: ClickHouse;

  constructor(opts: {
    clickhouseUrl: string;
  }) {
    this.clickhouse = new ClickHouse({ url: opts.clickhouseUrl });
  }

  public get insertSdkTelemetry() {
    return this.clickhouse.client.insert({
      table: "telemetry.raw_sdks_v1",
      schema: z.object({
        request_id: z.string(),
        time: z.number().int(),
        runtime: z.string(),
        platform: z.string(),
        versions: z.array(z.string()),
      }),
    });
  }

  public get insertRatelimit() {
    return this.clickhouse.ratelimits.insert;
  }

  public get insertKeyVerification() {
    return this.clickhouse.verifications.insert;
  }

  public get insertApiRequest() {
    return this.clickhouse.api.insert;
  }

  public get getVerificationsDaily() {
    return this.clickhouse.verifications.perDay;
  }
}