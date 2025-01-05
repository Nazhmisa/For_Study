class Task:
    def __init__(self, title, description, priority):
        self.title = title
        self.description = description
        self.priority = priority
    
    def __str__(self):
        task_print = f"Приоритет: {self.priority} \nЗадача: {self.title} \nОписание: {self.description}"
        return task_print


list_of_tasks = []

def add_task():
    title_of_task = input("Введите название задачи: ")
    desc_of_task = input("Введите описание задачи: ")
    prior_of_task = input("Введите приоритет задачи: ")
    new_task = Task(title_of_task, desc_of_task, prior_of_task)
    list_of_tasks.append(new_task)
    return new_task

def delete_task():
    num_of_task = int(input("Введите номер задачи, которую хотите удалить: "))
    if num_of_task < 1 or num_of_task > len(list_of_tasks):
        print("Введенное значение за пределами количества задач")
    else:
        num_of_task -= 1
        list_of_tasks.pop(num_of_task)
        print('Задача успешно удалена!')


def choice_for_user():
    num_for_add = int(input("Вы хотите добавить, посмотреть или удалить задачу ? (1 - чтобы посмотреть, 2 - чтобы добавить, 3 - чтобы удалить): "))
    if num_for_add == 1:
        if len(list_of_tasks) == 0:
            print('Задач пока нет, если хотите добавить, нажмите 2')
        else:
            for task in list_of_tasks:
                print(task)
    elif num_for_add == 2:
        add_task()
    elif num_for_add == 3:
        if len(list_of_tasks) == 0:
            print("Нет задач для удаления.")
        else:
            delete_task()
    else:
        print("Такого выбора нет, попробуйте еще раз.")

    choice_for_user()


choice_for_user()

