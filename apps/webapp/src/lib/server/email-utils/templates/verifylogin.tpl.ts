/**
 * Invitation email template options
 */
export interface Options {
	link: string;
}

/**
 * Invitation email subject template
 */
export function subject() {
	return `You have requested a link to sign-in to Provar.se`;
}

/**
 * Invitation email text template
 */
export function text(options: Options) {
	const { link } = options;
	return `Hello,
You have requested a link to sign-in to Provar.se. To access your account, click the link below:

${link}

If you do not wish to join this organization you can safely ignore it.

Thanks,
The Provar Team`;
}

/**
 * Invitation email html template
 */
export function html(options: Options) {
	const { link } = options;
	return `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <title>You have requested a link to sign-in to Provar.se</title>
</head>
<body style="margin: 0; padding: 0; font-family: 'Segoe UI', sans-serif; background-color: #f8fafc;">
  <table width="100%" border="0" cellspacing="0" cellpadding="0" bgcolor="#f8fafc">
    <tr>
      <td align="center" style="padding: 50px 0;">
        <a href="https://provar.se">
          <img alt="Provar" height="64" src="https://storage.googleapis.com/provar-assets/email/logo-64.png" style="max-width: 100%; vertical-align: middle; line-height: 1">
        </a>
      </td>
    </tr>
    <tr>
      <td align="center">
        <table width="100%" border="0" cellspacing="0" cellpadding="0" style="max-width: 600px;">
          <tr>
            <td bgcolor="#ffffff" style="padding: 30px; border-radius: 8px; box-shadow: 0 2px 4px rgba(0,0,0,0.1);">
              <h1 style="margin: 0 0 20px; font-size: 24px; font-weight: 600;">Hello,</h1>
              <p style="margin: 0 0 20px; font-size: 16px; line-height: 1.5;">You have requested a link to sign in to Provar.se. To access your account, click the link below:</p>
              <a href="${link}" style="display: inline-block; padding: 12px 24px; background-color: #4338ca; color: #ffffff; text-decoration: none; font-weight: 600; border-radius: 4px;">Sign in to Provar.se →</a>
              <p style="margin: 20px 0 0; font-size: 16px; line-height: 1.5;">If you did not request this email, you can safely ignore it.</p>
              <hr style="margin: 32px 0; border: none; height: 1px; background-color: #e2e8f0;">
              <p style="margin: 0; font-size: 16px;">Thanks,<br>The Provar Team</p>
            </td>
          </tr>
        </table>
      </td>
    </tr>
    <tr>
      <td align="center" style="padding: 50px 0;"><br></td>
    </tr>
  </table>
</body>
</html>
`;
}
