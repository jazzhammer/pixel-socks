const { defineConfig } = require("cypress");

module.exports = defineConfig({
  projectId: "14qywa",
  e2e: {
    setupNodeEvents(on, config) {
      // implement node event listeners here
    },
    "excludeSpecPattern": "**/examples/*",
    "specPattern": "cypress/integration/pixel/**/*.spec.js"
  },
});
