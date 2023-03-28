import { chromium } from 'k6/x/browser';

export default async function () {
  const browser = chromium.launch({
    headless: __ENV.XK6_HEADLESS ? true : false,
  });
  const page = browser.newPage();

  await page.goto('https://test.k6.io/my_messages.php', { waitUntil: 'networkidle' });
    
  const userInput = page.locator('input[name="login"]');
  await userInput.click();
  page.keyboard.type('admin');
    
  const pwdInput = page.locator('input[name="password"]');
  await pwdInput.click();
  page.keyboard.type('123');

  page.keyboard.press('Enter'); // submit
    
  await page.close();
  await browser.close();
}