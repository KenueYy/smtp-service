package templates

import (
	"fmt"
	"time"
)

func SubscriptionExpiringSoon(expireDate time.Time, renewalURL string, daysLeft int) string {
	dateStr := expireDate.Format("02.01.2006")
	heading := subscriptionHeading(daysLeft)

	return fmt.Sprintf(`<!DOCTYPE html>
<html lang="ru">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<meta http-equiv="X-UA-Compatible" content="IE=edge">
<title>neVPN — Подписка заканчивается</title>
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

<!-- Clock Icon -->
<tr>
<td align="center" style="padding:8px 24px 28px;">
<table role="presentation" cellpadding="0" cellspacing="0" border="0">
<tr>
<td style="width:56px;height:56px;background-color:#fef3c7;border-radius:50%%;text-align:center;vertical-align:middle;">
<span style="font-size:28px;line-height:56px;">⏰</span>
</td>
</tr>
</table>
</td>
</tr>

<!-- Heading -->
<tr>
<td align="center" style="padding:0 32px 8px;">
<h1 style="margin:0;font-family:'Inter',-apple-system,BlinkMacSystemFont,'Segoe UI',Roboto,Helvetica,Arial,sans-serif;font-size:22px;font-weight:700;color:#0a1929;line-height:1.3;">
%s
</h1>
</td>
</tr>

<!-- Subtitle -->
<tr>
<td align="center" style="padding:0 32px 24px;">
<p style="margin:0;font-family:'Inter',-apple-system,BlinkMacSystemFont,'Segoe UI',Roboto,Helvetica,Arial,sans-serif;font-size:14px;color:#486581;line-height:1.6;">
Дата окончания: <span style="color:#b45309;font-weight:600;">%s</span>
</p>
</td>
</tr>

<!-- CTA -->
<tr>
<td align="center" style="padding:0 32px 36px;">
<a href="%s" style="display:inline-block;background-color:#102a43;color:#ffffff;text-decoration:none;padding:14px 36px;border-radius:10px;font-family:'Inter',-apple-system,BlinkMacSystemFont,'Segoe UI',Roboto,Helvetica,Arial,sans-serif;font-size:15px;font-weight:600;line-height:1.2;">
Продлить подписку
</a>
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
После окончания подписки доступ к VPN будет <span style="color:#486581;font-weight:500;">ограничен</span>.
<br>
Продлите подписку сейчас, чтобы не потерять доступ.
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
</html>`, heading, dateStr, renewalURL)
}

func subscriptionHeading(daysLeft int) string {
	switch {
	case daysLeft >= 3:
		return "Подписка заканчивается через 3 дня"
	case daysLeft == 2:
		return "Подписка заканчивается через 2 дня"
	case daysLeft == 1:
		return "Подписка заканчивается завтра"
	default:
		return "Подписка заканчивается сегодня"
	}
}

func SubscriptionExpired(expireDate time.Time, renewalURL string) string {
	dateStr := expireDate.Format("02.01.2006")
	return fmt.Sprintf(`<!DOCTYPE html>
<html lang="ru">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<meta http-equiv="X-UA-Compatible" content="IE=edge">
<title>neVPN — Подписка закончилась</title>
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

<!-- Warning Icon -->
<tr>
<td align="center" style="padding:8px 24px 28px;">
<table role="presentation" cellpadding="0" cellspacing="0" border="0">
<tr>
<td style="width:56px;height:56px;background-color:#fee2e2;border-radius:50%%;text-align:center;vertical-align:middle;">
<span style="font-size:28px;line-height:56px;">⚠️</span>
</td>
</tr>
</table>
</td>
</tr>

<!-- Heading -->
<tr>
<td align="center" style="padding:0 32px 8px;">
<h1 style="margin:0;font-family:'Inter',-apple-system,BlinkMacSystemFont,'Segoe UI',Roboto,Helvetica,Arial,sans-serif;font-size:22px;font-weight:700;color:#0a1929;line-height:1.3;">
Подписка уже закончилась
</h1>
</td>
</tr>

<!-- Subtitle -->
<tr>
<td align="center" style="padding:0 32px 24px;">
<p style="margin:0;font-family:'Inter',-apple-system,BlinkMacSystemFont,'Segoe UI',Roboto,Helvetica,Arial,sans-serif;font-size:14px;color:#486581;line-height:1.6;">
Дата окончания: <span style="color:#dc2626;font-weight:600;">%s</span>
</p>
</td>
</tr>

<!-- Warning text -->
<tr>
<td align="center" style="padding:0 32px 24px;">
<p style="margin:0;font-family:'Inter',-apple-system,BlinkMacSystemFont,'Segoe UI',Roboto,Helvetica,Arial,sans-serif;font-size:14px;color:#486581;line-height:1.6;">
Доступ к VPN может быть <span style="color:#dc2626;font-weight:600;">ограничен</span>.
</p>
</td>
</tr>

<!-- CTA -->
<tr>
<td align="center" style="padding:0 32px 36px;">
<a href="%s" style="display:inline-block;background-color:#102a43;color:#ffffff;text-decoration:none;padding:14px 36px;border-radius:10px;font-family:'Inter',-apple-system,BlinkMacSystemFont,'Segoe UI',Roboto,Helvetica,Arial,sans-serif;font-size:15px;font-weight:600;line-height:1.2;">
Продлить подписку
</a>
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
Для восстановления доступа продлите подписку в личном кабинете.
<br>
Если у вас есть вопросы — напишите в поддержку.
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
</html>`, dateStr, renewalURL)
}
