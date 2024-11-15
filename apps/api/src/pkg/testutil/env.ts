import { z } from "zod";

export const databaseEnv = z.object({
  DATABASE_HOST: z.string().default("localhost:3900"),
  DATABASE_USERNAME: z.string().default("ghost"),
  DATABASE_PASSWORD: z.string().default("password"),
});

export const integrationTestEnv = databaseEnv.merge(
  z.object({
    GHOST_BASE_URL: z.string().url().default("http://localhost:8787"),
  }),
);

export const benchmarkTestEnv = databaseEnv.merge(
  z.object({
    PLANETFALL_URL: z.string().url(),
    PLANETFALL_API_KEY: z.string(),
    GHOST_BASE_URL: z.string().url(),
  }),
);
