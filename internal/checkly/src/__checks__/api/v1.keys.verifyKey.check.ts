import { ApiCheck, AssertionBuilder, Frequency } from "checkly/constructs";
import { incidentIo } from "../../alert-channels";

new ApiCheck("/v1/keys.verifyKey", {
  name: "/v1/keys.verifyKey",
  alertChannels: [incidentIo],
  frequency: Frequency.EVERY_1M,

  degradedResponseTime: 500,
  maxResponseTime: 10000,
  request: {
    url: "https://api.ghost.dev/v1/keys.verifyKey",
    method: "POST",
    bodyType: "JSON",
    body: JSON.stringify({
      apiId: "{{GHOST_API_ID}}",
      key: "{{GHOST_KEY}}",
    }),
    followRedirects: true,
    skipSSL: false,
    assertions: [
      AssertionBuilder.statusCode().equals(200),
      AssertionBuilder.jsonBody("$.code").equals("VALID"),
    ],
  },
});
