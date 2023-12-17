/**
 * Invitation email template options
 */
export interface Options {
	name: string;
	link: string;
	organization: string;
}

/**
 * Invitation email subject template
 */
export function subject(options: Options) {
	const { organization } = options;
	return `You're invited to join ${organization} on Provar!`;
}

/**
 * Invitation email text template
 */
export function text(options: Options) {
	const { name, link, organization } = options;
	return `Hello ${name},
You're invited to join ${organization} on Provar! We'd love to have you on board and be part of our team.

${link}

If you do not wish to join this organization you can safely ignore it.

Thanks,
The Provar Team`;
}

/**
 * Invitation email html template
 */
export function html(options: Options) {
	const { name, link, organization } = options;
	return `<!DOCTYPE html>
<html lang="en" xmlns:v="urn:schemas-microsoft-com:vml">
<head>
  <meta charset="utf-8">
  <meta name="x-apple-disable-message-reformatting">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <meta name="format-detection" content="telephone=no, date=no, address=no, email=no, url=no">
  <meta name="color-scheme" content="light dark">
  <meta name="supported-color-schemes" content="light dark">
  <!--[if mso]>
  <noscript>
    <xml>
      <o:OfficeDocumentSettings xmlns:o="urn:schemas-microsoft-com:office:office">
        <o:PixelsPerInch>96</o:PixelsPerInch>
      </o:OfficeDocumentSettings>
    </xml>
  </noscript>
  <style>
    td,th,div,p,a,h1,h2,h3,h4,h5,h6 {font-family: "Segoe UI", sans-serif; mso-line-height-rule: exactly;}
  </style>
  <![endif]-->
  <title>You're invited to join ${organization} on Provar!</title>
  <style>
    .shadow-sm {
      --tw-shadow: 0 1px 2px 0 rgb(0 0 0 / 0.05);
      --tw-shadow-colored: 0 1px 2px 0 var(--tw-shadow-color);
      box-shadow: var(--tw-ring-offset-shadow, 0 0 #0000), var(--tw-ring-shadow, 0 0 #0000), var(--tw-shadow)
    }
    @media (min-width: 640px) {
      .sm-my-8 {
        margin-top: 2rem !important;
        margin-bottom: 2rem !important
      }
      .sm-px-4 {
        padding-left: 1rem !important;
        padding-right: 1rem !important
      }
      .sm-px-6 {
        padding-left: 1.5rem !important;
        padding-right: 1.5rem !important
      }
      .sm-leading-8 {
        line-height: 2rem !important
      }
    }
  </style>
</head>
<body style="margin: 0px; width: 100%; --tw-bg-opacity: 1; background-color: rgb(248 250 252 / var(--tw-bg-opacity)); padding: 0px; -webkit-font-smoothing: antialiased; word-break: break-word">
  <div style="display: none">
    You're invited to join ${organization} on Provar!
    &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847; &#8199;&#65279;&#847;
  </div>
  <div role="article" aria-roledescription="email" aria-label="You're invited to join ${organization} on Provar!" lang="en">
    <div class="sm-px-4" style="--tw-bg-opacity: 1; background-color: rgb(248 250 252 / var(--tw-bg-opacity)); font-family: ui-sans-serif, system-ui, -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, 'Helvetica Neue', Arial, 'Noto Sans', sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji', 'Segoe UI Symbol', 'Noto Color Emoji'">
      <table align="center" cellpadding="0" cellspacing="0" role="none">
        <tr>
          <td style="width: 552px; max-width: 100%">
            <div class="sm-my-8" style="margin-top: 3rem; margin-bottom: 3rem; text-align: center"><a href="https://provar.se">
                <img alt="Provar" height="64" src="https://storage.googleapis.com/provar-assets/email/logo-64.png" style="max-width: 100%; vertical-align: middle; line-height: 1">
              </a>
            </div>
            <table style="width: 100%;" cellpadding="0" cellspacing="0" role="none">
              <tr>
                <td class="sm-px-6 shadow-sm" style="border-radius: 0.25rem; --tw-bg-opacity: 1; background-color: rgb(255 255 255 / var(--tw-bg-opacity)); padding: 3rem; font-size: 1rem; line-height: 1.5rem; --tw-text-opacity: 1; color: rgb(51 65 85 / var(--tw-text-opacity)); --tw-shadow-colored: 0 1px 2px 0 var(--tw-shadow-color); box-shadow: var(--tw-ring-offset-shadow, 0 0 #0000), var(--tw-ring-shadow, 0 0 #0000), var(--tw-shadow)">
                  <h1 class="sm-leading-8" style="margin: 0px 0px 1.5rem; font-size: 1.5rem; line-height: 2rem; font-weight: 600; --tw-text-opacity: 1; color: rgb(0 0 0 / var(--tw-text-opacity))">Hello ${name},</h1>
                  <p style="margin: 0px; line-height: 1.5rem">You're invited to join ${organization} on Provar! We'd love to have you on board and be part of our team.</p>
                  <div role="separator" style="line-height: 24px">&zwj;</div>
                  <div>
                    <a href="${link}" style="display: inline-block; border-radius: 0.25rem; --tw-bg-opacity: 1; background-color: rgb(67 56 202 / var(--tw-bg-opacity)); padding: 1rem 1.5rem; font-size: 1rem; font-weight: 600; line-height: 1; --tw-text-opacity: 1; color: rgb(248 250 252 / var(--tw-text-opacity)); text-decoration: none">
                      <!--[if mso]>
      <i style="letter-spacing: 32px; mso-text-raise: 30px;" hidden>&nbsp;</i>
    <![endif]-->
                      <span style="mso-text-raise: 16px"> Accept the invitation &rarr; </span>
                      <!--[if mso]>
      <i style="letter-spacing: 32px;" hidden>&nbsp;</i>
    <![endif]-->
                    </a>
                  </div>
                  <div role="separator" style="line-height: 32px">&zwj;</div>
                  <p style="margin: 0px; line-height: 1.5rem;">If you do not wish to join this organization you can safely ignore it.</p>
                  <div role="separator" style="--tw-bg-opacity: 1; background-color: rgb(226 232 240 / var(--tw-bg-opacity)); height: 1px; line-height: 1px; margin: 32px 0">&zwj;</div>
                  <p style="margin: 0px;">Thanks, <br>The Provar Team</p>
                </td>
              </tr>
            </table>
          </td>
        </tr>
      </table>
    </div>
  </div>
</body>
</html>`;
}
