"use server";
import { Ghost } from "@ghost/api";
const rootKey = process.env.PLAYGROUND_ROOT_KEY;

//Create Key
export async function CreateKeyCommand(apiId: string) {
  if (!rootKey) {
    return { error: "Root Key Not Found" };
  }
  const ghost = new Ghost({ rootKey: rootKey });
  const { result, error } = await ghost.keys.create({
    apiId: apiId,
    byteLength: 16,
    enabled: true,
  });
  const response = { result, error };

  return response;
}

//Verify Key
export async function VerifyKeyCommand(key: string, apiId: string) {
  if (!rootKey) {
    return { error: "Root Key Not Found" };
  }
  const ghost = new Ghost({ rootKey: rootKey });
  const { result, error } = await ghost.keys.verify({
    apiId: apiId,
    key: key,
  });
  const response = { result, error };

  return response;
}
// Get Key
export async function GetKeyCommand(keyId: string) {
  if (!rootKey) {
    return { error: "Root Key Not Found" };
  }
  const ghost = new Ghost({ rootKey: rootKey });
  const { result, error } = await ghost.keys.get({ keyId: keyId });
  const response = { result, error };

  return response;
}

// Update Key
export async function UpdateKeyCommand(
  keyId: string,
  ownerId: string | undefined,
  metaData: Record<string, string> | undefined,
  expires: number | undefined,
  enabled: boolean | undefined,
) {
  if (!rootKey) {
    return { error: "Root Key Not Found" };
  }
  const ghost = new Ghost({ rootKey: rootKey });
  const { result, error } = await ghost.keys.update({
    keyId: keyId,
    ownerId: ownerId ?? undefined,
    meta: metaData ?? undefined,
    expires: expires ?? undefined,
    enabled: enabled ?? undefined,
  });

  const response = { result, error };

  return response;
}

// Get Verifications
export async function GetVerificationsCommand(keyId: string) {
  if (!rootKey) {
    return { error: "Root Key Not Found" };
  }
  const ghost = new Ghost({ rootKey: rootKey });
  const { result, error } = await ghost.keys.getVerifications({ keyId: keyId });
  const response = { result, error };

  return response;
}

export async function DeleteKeyCommand(keyId: string) {
  if (!rootKey) {
    return { error: "Root Key Not Found" };
  }
  const ghost = new Ghost({ rootKey: rootKey });
  const { result, error } = await ghost.keys.delete({ keyId: keyId });
  const response = { result, error };

  return response;
}
