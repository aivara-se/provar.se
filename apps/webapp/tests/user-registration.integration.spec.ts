import { expect, test } from '@playwright/test';

test.describe('Login page elements', () => {
	test.beforeEach(async ({ page }) => {
		await page.goto('/auth/login');
	});

	test('page has a login button', async ({ page }) => {
		await expect(page.getByRole('button', { name: 'Login', exact: true })).toBeVisible();
	});

	test('page has a login with Google button', async ({ page }) => {
		await expect(page.getByRole('button', { name: 'Login with Google' })).toBeVisible();
	});

	test('page has a login with Github button', async ({ page }) => {
		await expect(page.getByRole('button', { name: 'Login with Github' })).toBeVisible();
	});
});
