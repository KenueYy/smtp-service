package templates

import (
	"fmt"
)

func AuthCodeEmail(code string) string {
	return fmt.Sprintf(`<!DOCTYPE html>
<html lang="ru">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<meta http-equiv="X-UA-Compatible" content="IE=edge">
<title>neVPN — Код подтверждения</title>
</head>
<body style="margin:0;padding:0;background-color:#f8fafc;-webkit-text-size-adjust:100%%;-ms-text-size-adjust:100%%;">
<table role="presentation" cellpadding="0" cellspacing="0" border="0" width="100%%" style="background-color:#f8fafc;">
<tr>
<td align="center" style="padding:48px 16px 60px;">

<!-- Main Card -->
<table role="presentation" cellpadding="0" cellspacing="0" border="0" width="100%%" style="max-width:440px;background-color:#ffffff;border-radius:12px;border:1px solid #e2e8f0;overflow:hidden;">

<!-- Logo -->
<tr>
<td align="center" style="padding:40px 24px 16px;">
<table role="presentation" cellpadding="0" cellspacing="0" border="0">
<tr>
<td style="font-family:'Inter',-apple-system,BlinkMacSystemFont,'Segoe UI',Roboto,Helvetica,Arial,sans-serif;font-size:26px;font-weight:700;color:#0a1929;letter-spacing:-0.5px;">
ne<span style="color:#486581;">VPN</span>
</td>
</tr>
</table>
</td>
</tr>

<!-- Tagline -->
<tr>
<td align="center" style="padding:0 24px 24px;">
<p style="margin:0;font-family:'Inter',-apple-system,BlinkMacSystemFont,'Segoe UI',Roboto,Helvetica,Arial,sans-serif;font-size:12px;color:#627d98;text-transform:uppercase;letter-spacing:1.5px;font-weight:500;">
Ускоритель интернета
</p>
</td>
</tr>

<!-- Shield Icon -->
<tr>
<td align="center" style="padding:8px 24px 28px;">
<table role="presentation" cellpadding="0" cellspacing="0" border="0">
<tr>
<td style="width:56px;height:56px;background-color:#e0e7ff;border-radius:50%%;text-align:center;vertical-align:middle;">
<span style="font-size:28px;line-height:56px;">🛡️</span>
</td>
</tr>
</table>
</td>
</tr>

<!-- Heading -->
<tr>
<td align="center" style="padding:0 32px 8px;">
<h1 style="margin:0;font-family:'Inter',-apple-system,BlinkMacSystemFont,'Segoe UI',Roboto,Helvetica,Arial,sans-serif;font-size:22px;font-weight:700;color:#0a1929;line-height:1.3;">
Код подтверждения
</h1>
</td>
</tr>

<!-- Subtitle -->
<tr>
<td align="center" style="padding:0 32px 28px;">
<p style="margin:0;font-family:'Inter',-apple-system,BlinkMacSystemFont,'Segoe UI',Roboto,Helvetica,Arial,sans-serif;font-size:14px;color:#486581;line-height:1.6;">
Используйте этот код для входа в аккаунт neVPN
</p>
</td>
</tr>

<!-- Code Box -->
<tr>
<td align="center" style="padding:0 32px 36px;">
<table role="presentation" cellpadding="0" cellspacing="0" border="0" style="width:100%%;">
<tr>
<td style="background-color:#f1f5f9;border:1px solid #e2e8f0;border-radius:10px;padding:20px 16px;text-align:center;">
<span style="font-family:'SF Mono','Fira Code','Cascadia Code',Consolas,monospace;font-size:36px;font-weight:700;color:#0a1929;letter-spacing:10px;line-height:1.2;">%s</span>
</td>
</tr>
</table>
</td>
</tr>

<!-- Divider -->
<tr>
<td style="padding:0 32px;">
<table role="presentation" cellpadding="0" cellspacing="0" border="0" width="100%%">
<tr>
<td style="border-top:1px solid #e2e8f0;height:1px;"></td>
</tr>
</table>
</td>
</tr>

<!-- Info -->
<tr>
<td align="center" style="padding:24px 32px 32px;">
<p style="margin:0;font-family:'Inter',-apple-system,BlinkMacSystemFont,'Segoe UI',Roboto,Helvetica,Arial,sans-serif;font-size:12px;color:#627d98;line-height:1.6;">
Код действителен в течение <span style="color:#486581;font-weight:500;">5 минут</span>.
<br>
Если вы не запрашивали код — просто проигнорируйте это письмо.
</p>
</td>
</tr>
</table>

<!-- Footer -->
<table role="presentation" cellpadding="0" cellspacing="0" border="0" width="100%%" style="max-width:440px;">
<tr>
<td align="center" style="padding:24px 24px;">
<p style="margin:0;font-family:'Inter',-apple-system,BlinkMacSystemFont,'Segoe UI',Roboto,Helvetica,Arial,sans-serif;font-size:11px;color:#9fb3c8;line-height:1.5;">
neVPN — ускоритель интернета, который просто работает
</p>
</td>
</tr>
</table>

</td>
</tr>
</table>
</body>
</html>`, code)
}
