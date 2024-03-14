import { expect, test } from '@playwright/test';

test.describe('Login', () => {
	test('the page has a buttons to login with email and with social auth', async ({ page }) => {
		await page.goto('/');
		await expect(page.getByText('Login with email')).toBeVisible();
		await expect(page.getByText('Login with Google')).toBeVisible();
		await expect(page.getByText('Login with Github')).toBeVisible();
	});
});
