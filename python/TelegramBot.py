from telegram import Update
from telegram.ext import ApplicationBuilder, CommandHandler, ContextTypes

async def start(update: Update, context: ContextTypes.DEFAULT_TYPE):
    await update.message.reply_text("""Привет! Я телеграмм-бот для ведения задач!
                                        Чтобы узнать больше о функционале бота, нажми /help
                                    """)

async def help(update: Update, context: ContextTypes.DEFAULT_TYPE):
    await update.message.reply_text("""У данного бота имеются следующие функции:
                                        /add_task - добавление задачи
                                        /task_list - вывод списка задач
                                        /delete_task - удалить задачу
                                    """)


if __name__ == "__main__":
    app = ApplicationBuilder().token("7843152834:AAGRDZQ_BTmUnyuDnG6yDUnKhazurtFhoE4").build()
    app.add_handler(CommandHandler("start", start))
    app.add_handler(CommandHandler("help", help))
    app.run_polling()