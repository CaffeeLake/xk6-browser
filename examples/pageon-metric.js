import { browser } from 'k6/x/browser/async';

export const options = {
  scenarios: {
    ui: {
      executor: 'shared-iterations',
      options: {
        browser: {
            type: 'chromium',
        },
      },
    },
  },
}

export default async function() {
  const page = await browser.newPage();

  page.on('metric', (metric) => {
    metric.Tag({
      urls: [
        {url: /^https:\/\/test\.k6\.io\/\?q=[0-9a-z]+$/, name:'test'},
      ]
    });
  });

  try {
    await page.goto('https://test.k6.io/?q=abc123');
    await page.goto('https://test.k6.io/?q=def456');
  } finally {
    await page.close();
  }
}