import path from "node:path";

import fs from "node:fs";
import * as clack from "@clack/prompts";
import { marshalEnv } from "../env";

const appPath = path.join(__dirname, "../../../../apps/www");
const envPath = path.join(appPath, ".env");

export function bootstrapWWW(resources: {
  workspace: { id: string };
  api: { id: string };
  webhooksApi: { id: string };
}) {
  const env = marshalEnv({
    Database: {
      DATABASE_HOST: "localhost:3900",
      DATABASE_USERNAME: "ghost",
      DATABASE_PASSWORD: "password",
    },
    Bootstrap: {
      GHOST_WORKSPACE_ID: resources.workspace.id,
      GHOST_API_ID: resources.api.id,
      GHOST_WEBHOOK_KEYS_API_ID: resources.webhooksApi.id,
    },
  });

  if (fs.existsSync(envPath)) {
    clack.log.warn(".env already exists, please add the variables manually");
    clack.note(env, envPath);
  } else {
    fs.writeFileSync(envPath, env);
    clack.log.step(`Wrote variables to ${envPath}`);
  }
}
