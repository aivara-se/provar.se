import { expect, test } from '@playwright/test';

test.describe('Login', () => {
	test('the page has a buttons to login with email and with social auth', async ({ page }) => {
		await page.goto('/');
		await expect(page.getByText('Login with email')).toBeVisible();
		await expect(page.getByText('Login with Google')).toBeVisible();
		await expect(page.getByText('Login with Github')).toBeVisible();
	});

	test('user is redirected to a page with instructions for email login', async ({ page }) => {
		await page.goto('/');
		await page.fill('input[name="email"]', 'testuser@provar.se');
		await page.click('text=Login with email');
		await expect(page.getByText('You will receive an email soon.')).toBeVisible();
	});
});
