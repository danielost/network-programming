import logging
import random
from telegram import (ReplyKeyboardMarkup, ReplyKeyboardRemove, Update,
                      InlineKeyboardButton, InlineKeyboardMarkup)
from telegram.ext import (Application, CallbackQueryHandler, CommandHandler,
                          ContextTypes, ConversationHandler, MessageHandler, filters)

logging.basicConfig(format='%(asctime)s - %(name)s - %(levelname)s - %(message)s',
                    level=logging.INFO)

logger = logging.getLogger(__name__)

LEAGUE, SEASON, TEAM_A, TEAM_B, SUMMARY = range(5)


async def start(update: Update, context: ContextTypes.DEFAULT_TYPE) -> int:
    reply_keyboard = [['Premier League', 'La Liga', 'Bundesliga']]

    await update.message.reply_text(
        '<b>Лабораторна робота 4 з ПвМС (варіант 6)\n'
        'Оберіть бажану лігу</b>',
        parse_mode='HTML',
        reply_markup=ReplyKeyboardMarkup(reply_keyboard, one_time_keyboard=True, resize_keyboard=True),
    )

    return LEAGUE


async def league(update: Update, context: ContextTypes.DEFAULT_TYPE) -> int:
    user = update.message.from_user
    context.user_data['league'] = update.message.text
    logger.info('League of %s: %s', user.first_name, update.message.text)
    await update.message.reply_text(
        f'<b>Ви обрали лігу {update.message.text}.\n'
        f'Далі оберіть сезон</b>',
        parse_mode='HTML',
        reply_markup=ReplyKeyboardRemove(),
    )

    keyboard = [
        [InlineKeyboardButton('2021', callback_data='2021')],
        [InlineKeyboardButton('2022', callback_data='2022')],
        [InlineKeyboardButton('2023', callback_data='2023')],
        [InlineKeyboardButton('2024', callback_data='2023')],
    ]
    reply_markup = InlineKeyboardMarkup(keyboard)
    await update.message.reply_text('<b>Будь ласка, оберіть:</b>', parse_mode='HTML', reply_markup=reply_markup)

    return SEASON


async def season(update: Update, context: ContextTypes.DEFAULT_TYPE) -> int:
    query = update.callback_query
    await query.answer()
    context.user_data['season'] = query.data
    await query.edit_message_text(
        text=f'<b>Ви обрали сезон {query.data}.\n'
             f'Далі оберіть першу команду.</b>',
        parse_mode='HTML'
    )

    keyboard = [
        [InlineKeyboardButton('Monaco', callback_data='Monaco')],
        [InlineKeyboardButton('Reims', callback_data='Reims')],
        [InlineKeyboardButton('Nice', callback_data='Nice')],
        [InlineKeyboardButton('Nantes', callback_data='Nantes')],
    ]
    reply_markup = InlineKeyboardMarkup(keyboard)
    await query.message.reply_text('<b>Оберіть команду:</b>', parse_mode='HTML', reply_markup=reply_markup)

    return TEAM_A


async def team_a(update: Update, context: ContextTypes.DEFAULT_TYPE) -> int:
    query = update.callback_query
    await query.answer()
    context.user_data['team_a'] = query.data
    await query.edit_message_text(
        text=f'<b>Ви обрали команду {query.data}.\n'
             f'Далі оберіть другу команду</b>',
        parse_mode='HTML'
    )

    keyboard = [
        [InlineKeyboardButton('Monaco', callback_data='Monaco')],
        [InlineKeyboardButton('Reims', callback_data='Reims')],
        [InlineKeyboardButton('Nice', callback_data='Nice')],
        [InlineKeyboardButton('Nantes', callback_data='Nantes')],
    ]
    reply_markup = InlineKeyboardMarkup(keyboard)
    await query.message.reply_text('<b>Оберіть команду:</b>', parse_mode='HTML', reply_markup=reply_markup)

    return TEAM_B


async def team_b(update: Update, context: ContextTypes.DEFAULT_TYPE) -> int:
    query = update.callback_query
    await query.answer()
    context.user_data['team_b'] = query.data
    await query.edit_message_text(
        text=f'<b>Ви обрали команду {query.data}.\n'
             f'Результат:</b>',
        parse_mode='HTML'
    )

    await summary(update, context)


async def summary(update: Update, context: ContextTypes.DEFAULT_TYPE) -> int:
    selections = context.user_data
    summary_text = (f"<b>Ви обрали:\n</b>"
                    f"<b>Ліга:</b> {selections.get('league')}\n"
                    f"<b>Сезон:</b> {selections.get('season')}\n"
                    f"<b>Команда А:</b> {selections.get('team_a')}\n"
                    f"<b>Команда Б:</b> {selections.get('team_b')}\n"
                    f"\n<b>Результат матчу: {random.randint(0, 5)}:{random.randint(0, 5)}</b>")

    chat_id = update.effective_chat.id

    await context.bot.send_message(chat_id=chat_id, text=summary_text, parse_mode='HTML')

    return ConversationHandler.END


async def cancel(update: Update, context: ContextTypes.DEFAULT_TYPE) -> int:
    """Cancels and ends the conversation."""
    await update.message.reply_text('Bye! Hope to talk to you again soon.', reply_markup=ReplyKeyboardRemove())
    return ConversationHandler.END


def main() -> None:
    """Run the bot."""
    application = Application.builder().token("6531150712:AAEli8RkCbJMGEde4_rl3GAqIJsW5TBajZk").build()

    conv_handler = ConversationHandler(
        entry_points=[CommandHandler('start', start)],
        states={
            LEAGUE: [MessageHandler(filters.TEXT & ~filters.COMMAND, league)],
            SEASON: [CallbackQueryHandler(season)],
            TEAM_A: [CallbackQueryHandler(team_a)],
            TEAM_B: [CallbackQueryHandler(team_b)],
            SUMMARY: [MessageHandler(filters.ALL, summary)]
        },
        fallbacks=[CommandHandler('cancel', cancel)],
    )

    application.add_handler(conv_handler)

    application.add_handler(CommandHandler('start', start))

    application.run_polling()


if __name__ == '__main__':
    main()