# Solace

## 1. Purpose, status, and authority of the document

This document is a self-contained product and technical specification for Solace. It intentionally covers all areas: product requirements, domain rules, API contracts, frontend behavior, backend architecture, runtime rules, and delivery rules - so that a developer can work using only this file.

## 2. Product context and MVP goals

Solace is an event planner with deep Telegram integration and an AI agent for planning in natural language. Telegram was chosen deliberately: the target audience is already there, the app is always at hand, notifications do not require installing a separate client, and they do not compete for trust with browser push notifications.

The goal of the MVP is to create a working, coherent prototype with clear backend boundaries, explicit REST contracts, and enough frontend structure to support all core scenarios. Full automated test coverage is not a priority at this stage. Manual smoke validation is mandatory for all significant changes.

All changes that affect behavior, architecture, runtime configuration, persistence, visual layout, or API contracts must update this document.

## 3. Roles and surfaces

The system provides two roles: user and administrator. Public visitors see only the landing page and the login page.

A user is any authenticated person using Solace to manage their events. The user works through the Telegram bot or the web interface. The user account is tightly bound to their Telegram ID.

An administrator is a role with access to admin tools. Administrators manage user limits, AI model configuration, and system settings. Administrators are assigned through runtime configuration or directly in the database.

## 4. Authentication and accounts

In the MVP, the only login method is Telegram OAuth. Email and password are not supported. An account is created on first authorization through the Telegram Login Widget on the website or on first interaction with the Telegram bot.

The user's identifier is their Telegram ID. When the account is created, the following are taken from the Telegram profile: name, username (if present), and language locale. The timezone is taken from the Telegram profile when available and can be overridden by the user in settings.

Future versions are planned to decouple the system from Telegram: Telegram will remain one of the login methods rather than the only one. This will require identifier migration. The current architecture must account for this and must not be built in a way that makes migration impossible.

## 5. User settings

Settings are a separate entity, available through the `/settings` command in the bot and through a modal window in the web interface. In the MVP, settings contain the timezone. Changing the timezone affects all future notifications and time display; past events are not recalculated.

## 6. Event entity

An event belongs to one user. Event fields:

- `title` - required text title of the event.
- `datetime` - required date and time of the event in UTC. Stored in UTC, displayed according to the user's timezone.
- `notes` - optional text description or notes.
- `recurrence` - optional recurrence settings. Format: RRULE (iCalendar standard). If recurrence is not specified, the event is a one-time event.
- `reminder_offset_minutes` - optional offset in minutes before the event when the reminder is sent. If not specified, no reminder is sent.

A recurring event physically exists as a single record with recurrence settings. Editing an event with recurrence applies to all future occurrences - there is no split into separate instances in the MVP. Deleting an event deletes it entirely together with all future occurrences.

Future versions are planned to support multiple reminders per event.

## 7. Notifications

Notifications are sent exclusively through the Telegram bot. For each event with `reminder_offset_minutes` set, a preliminary reminder is sent the specified number of minutes before the event. When the event time arrives, a separate notification is sent.

Notifications are tied to the user's Telegram account. If the user has blocked the bot, the notification is not delivered; the system logs the error but is not blocked.

## 8. Telegram bot

The bot is the main interaction channel for mobile users. It supports two modes: commands for direct event management and free-text input for the AI agent.

Standard commands in the MVP:

- `/start` - account initialization, welcome message.
- `/new` - create a new event through a step-by-step dialog.
- `/list` - list of all upcoming events for the user.
- `/today` - events for the user's current day (according to their timezone).
- `/settings` - view and change account settings.
- `/help` - list of available commands with a short description.

Any non-command message is passed to the AI agent (see section 10).

## 9. Web interface

The website provides a full user interface for working with events. Login is exclusively through Telegram OAuth (Telegram Login Widget).

In the MVP, the web interface contains the following sections:

Event list - a chronological list of all upcoming events for the user. Allows creating, editing, and deleting events through a form. The "calendar" view mode is postponed to post-MVP.

AI chat - an interface for creating events through the AI agent in free-text form.

Account settings - management of timezone and other user parameters.

Admin panel - available only to users with the administrator role. It is placed on a separate subdomain or in a separate navigation tab. Described in more detail in section 11.

## 10. AI agent

The AI agent allows the user to create events in natural language. The agent receives a text request, extracts the event parameters (`title`, `datetime`, `notes`, `recurrence`, `reminder_offset_minutes`), and creates the event. If there are not enough parameters to create the event, the agent asks clarifying questions.

In the MVP, the agent supports only event creation. CRUD operations through AI (editing, deletion, search) are postponed to post-MVP and will be implemented through MCP and additional agent tools.

Chat history with the agent is not stored between sessions (stateless). Within one session, dialog context is maintained to allow clarifications. Between sessions, the context is reset.

The agent runs on self-hosted Ollama instances. Limits are counted in tokens per user. Specific models and their parameters are managed through admin tools.

AI requests are subject to the rate limit policy. When the limit is exceeded, the user receives an informative error message indicating when the limit will be reset.

Future versions are planned to support BYOK (Bring Your Own Key) - the ability for a user to connect their own API key to Claude or other providers. This is a paid feature.

## 11. Admin tools

Administrators manage the system through two interfaces: Telegram commands for quick operations and a web panel (a separate tab or subdomain) for full management.

Admin tool capabilities in the MVP:

- Managing token and request limits per user (view, change, reset).
- Enabling and disabling AI models.
- Viewing AI usage statistics (tokens, requests, active users).
- Managing the administrator role (assign, revoke).

## 12. Rate limiting and AI limits

Each user has a token limit for using the AI agent. Limits are set globally by default and can be overridden for a specific user through admin tools. The limit type in the MVP is tokens because Ollama conveniently counts tokens.

When the limit is exceeded, the request is rejected at the backend level before reaching the model. The user receives a limit exceeded message. The logic for resetting limits (daily, monthly) is determined by configuration.

## 13. Roadmap (post-MVP)

The following capabilities are outside the scope of the MVP and reserved for future versions:

- Calendar view mode for events in the web interface.
- "Upcoming events" widget in the web interface.
- Multiple reminders for one event.
- AI full CRUD through MCP: editing, deleting, and searching events through the agent.
- BYOK - connecting user API keys to external AI providers (paid feature).
- Decoupling from Telegram: Telegram becomes one of the login methods, OIDC and alternative providers appear.
- Monetization: paid tiers, subscriptions, BYOK surcharge.
