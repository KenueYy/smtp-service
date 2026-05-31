package templates

import (
	"fmt"
	"html"
)

func SupportTicketEmail(description, contactMethod, contact string) string {
	return fmt.Sprintf(`<!DOCTYPE html>
<html lang="ru">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>neVPN — Новое обращение в поддержку</title>
</head>
<body style="margin:0;padding:0;background-color:#f8fafc;">
<table role="presentation" cellpadding="0" cellspacing="0" border="0" width="100%%" style="background-color:#f8fafc;">
<tr>
<td align="center" style="padding:48px 16px 60px;">

<table role="presentation" cellpadding="0" cellspacing="0" border="0" width="100%%" style="max-width:480px;background-color:#ffffff;border-radius:12px;border:1px solid #e2e8f0;overflow:hidden;">

<!-- Header -->
<tr>
<td style="padding:32px 32px 0;">
<p style="margin:0;font-family:'Inter',-apple-system,BlinkMacSystemFont,'Segoe UI',Roboto,sans-serif;font-size:12px;color:#627d98;text-transform:uppercase;letter-spacing:1.5px;font-weight:500;">
Новое обращение
</p>
</td>
</tr>

<!-- Icon -->
<tr>
<td align="center" style="padding:16px 32px 24px;">
<div style="width:48px;height:48px;background-color:#e0e7ff;border-radius:50%%;text-align:center;line-height:48px;font-size:24px;">💬</div>
</td>
</tr>

<!-- Description -->
<tr>
<td style="padding:0 32px 8px;">
<p style="margin:0;font-family:'Inter',-apple-system,BlinkMacSystemFont,'Segoe UI',Roboto,sans-serif;font-size:13px;font-weight:600;color:#0a1929;">
Описание проблемы
</p>
</td>
</tr>
<tr>
<td style="padding:0 32px 16px;">
<table role="presentation" cellpadding="0" cellspacing="0" border="0" width="100%%">
<tr>
<td style="background-color:#f1f5f9;border:1px solid #e2e8f0;border-radius:8px;padding:16px;">
<p style="margin:0;font-family:'Inter',-apple-system,BlinkMacSystemFont,'Segoe UI',Roboto,sans-serif;font-size:14px;color:#334155;line-height:1.6;white-space:pre-wrap;">%s</p>
</td>
</tr>
</table>
</td>
</tr>

<!-- Contact Method -->
<tr>
<td style="padding:0 32px 8px;">
<p style="margin:0;font-family:'Inter',-apple-system,BlinkMacSystemFont,'Segoe UI',Roboto,sans-serif;font-size:13px;font-weight:600;color:#0a1929;">
Способ связи
</p>
</td>
</tr>
<tr>
<td style="padding:0 32px 8px;">
<p style="margin:0;font-family:'Inter',-apple-system,BlinkMacSystemFont,'Segoe UI',Roboto,sans-serif;font-size:14px;color:#334155;">%s</p>
</td>
</tr>

<!-- Contact -->
<tr>
<td style="padding:0 32px 8px;">
<p style="margin:0;font-family:'Inter',-apple-system,BlinkMacSystemFont,'Segoe UI',Roboto,sans-serif;font-size:13px;font-weight:600;color:#0a1929;">
Контакт
</p>
</td>
</tr>
<tr>
<td style="padding:0 32px 24px;">
<p style="margin:0;font-family:'SF Mono','Fira Code','Cascadia Code',Consolas,monospace;font-size:14px;color:#0a1929;background-color:#f1f5f9;border-radius:6px;padding:8px 12px;display:inline-block;">%s</p>
</td>
</tr>

<!-- Divider -->
<tr>
<td style="padding:0 32px;">
<table role="presentation" cellpadding="0" cellspacing="0" border="0" width="100%%">
<tr><td style="border-top:1px solid #e2e8f0;height:1px;"></td></tr>
</table>
</td>
</tr>

<!-- Footer -->
<tr>
<td align="center" style="padding:24px 32px 32px;">
<p style="margin:0;font-family:'Inter',-apple-system,BlinkMacSystemFont,'Segoe UI',Roboto,sans-serif;font-size:11px;color:#9fb3c8;line-height:1.5;">
Автоматическое уведомление neVPN — support ticket
</p>
</td>
</tr>
</table>

</td>
</tr>
</table>
</body>
</html>`,
		html.EscapeString(description),
		html.EscapeString(contactMethod),
		html.EscapeString(contact),
	)
}
