package templates

// LoginEmail implements EmailTemplate
type LoginEmail struct{}

// SubjectText returns the subject of the email
func (t *LoginEmail) SubjectText(data map[string]string) string {
	return "You have requested a link to sign-in to Provar.se"
}

// ContentText returns the text content of the email
func (t *LoginEmail) ContentText(data map[string]string) string {
	return `Hello,
You have requested a link to sign-in to Provar.se. To access your account, click the link below:

` + data["link"] + `

If you do not wish to join this organization you can safely ignore it.

Thanks,
The Provar Team`
}

// ContentHTML returns the HTML content of the email
func (t *LoginEmail) ContentHTML(data map[string]string) string {
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
          <img alt="Provar" height="64" src="https://provar.se/assets/logo-512.png" style="max-width: 100%; vertical-align: middle; line-height: 1">
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
              <a href="` + data["link"] + `" style="display: inline-block; padding: 12px 24px; background-color: #4338ca; color: #ffffff; text-decoration: none; font-weight: 600; border-radius: 4px;">Sign in to Provar.se →</a>
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
`
}
