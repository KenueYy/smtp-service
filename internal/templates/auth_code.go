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
<title>неВПН — Код подтверждения</title>
</head>
<body style="margin:0;padding:0;background-color:#0a0a0f;-webkit-text-size-adjust:100%%;-ms-text-size-adjust:100%%;">
<table role="presentation" cellpadding="0" cellspacing="0" border="0" width="100%%" style="background-color:#0a0a0f;">
<tr>
<td align="center" style="padding:40px 16px 60px;">

<!-- Main Card -->
<table role="presentation" cellpadding="0" cellspacing="0" border="0" width="100%%" style="max-width:440px;background-color:#13131a;border-radius:16px;border:1px solid #1e1e2a;overflow:hidden;">
<!-- Logo -->
<tr>
<td align="center" style="padding:36px 24px 12px;">
<table role="presentation" cellpadding="0" cellspacing="0" border="0">
<tr>
<td style="font-family:-apple-system,BlinkMacSystemFont,'Segoe UI',Roboto,Helvetica,Arial,sans-serif;font-size:28px;font-weight:800;color:#ffffff;letter-spacing:-0.5px;">
не<span style="color:#6c5ce7;">ВПН</span>
</td>
</tr>
</table>
</td>
</tr>

<!-- Shield Icon -->
<tr>
<td align="center" style="padding:8px 24px 28px;">
<table role="presentation" cellpadding="0" cellspacing="0" border="0">
<tr>
<td style="width:56px;height:56px;background:linear-gradient(135deg,#6c5ce720,#a855f720);border-radius:50%%;text-align:center;vertical-align:middle;">
<span style="font-size:28px;line-height:56px;">🛡️</span>
</td>
</tr>
</table>
</td>
</tr>

<!-- Heading -->
<tr>
<td align="center" style="padding:0 32px 8px;">
<h1 style="margin:0;font-family:-apple-system,BlinkMacSystemFont,'Segoe UI',Roboto,Helvetica,Arial,sans-serif;font-size:22px;font-weight:700;color:#ffffff;line-height:1.3;">
Код подтверждения
</h1>
</td>
</tr>

<!-- Subtitle -->
<tr>
<td align="center" style="padding:0 32px 28px;">
<p style="margin:0;font-family:-apple-system,BlinkMacSystemFont,'Segoe UI',Roboto,Helvetica,Arial,sans-serif;font-size:14px;color:#8b8b9e;line-height:1.5;">
Используйте этот код для входа в аккаунт неВПН
</p>
</td>
</tr>

<!-- Code Box -->
<tr>
<td align="center" style="padding:0 32px 32px;">
<table role="presentation" cellpadding="0" cellspacing="0" border="0" style="width:100%%;">
<tr>
<td style="background:linear-gradient(135deg,#6c5ce715,#a855f715);border:1px solid #6c5ce730;border-radius:12px;padding:20px 16px;text-align:center;">
<span style="font-family:'SF Mono','Fira Code','Cascadia Code',Consolas,monospace;font-size:36px;font-weight:700;color:#ffffff;letter-spacing:10px;line-height:1.2;">%s</span>
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
<td style="border-top:1px solid #1e1e2a;height:1px;"></td>
</tr>
</table>
</td>
</tr>

<!-- Info -->
<tr>
<td align="center" style="padding:24px 32px 32px;">
<p style="margin:0;font-family:-apple-system,BlinkMacSystemFont,'Segoe UI',Roboto,Helvetica,Arial,sans-serif;font-size:12px;color:#5a5a72;line-height:1.6;">
Код действителен в течение <span style="color:#8b8b9e;font-weight:500;">5 минут</span>.
<br>
Если вы не запрашивали код — просто проигнорируйте это письмо.
</p>
</td>
</tr>
</table>

<!-- Footer -->
<table role="presentation" cellpadding="0" cellspacing="0" border="0" width="100%%" style="max-width:440px;">
<tr>
<td align="center" style="padding:20px 24px;">
<p style="margin:0;font-family:-apple-system,BlinkMacSystemFont,'Segoe UI',Roboto,Helvetica,Arial,sans-serif;font-size:11px;color:#3a3a4a;line-height:1.5;">
неВПН — безопасность без границ
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
