from tkinter import *
from tkinter import ttk
import sqlite3 as sql
from PIL import Image, ImageTk
import numpy as np
from tkinter.ttk import Combobox
import time



def db_roomstest(id, number, status):

    con = sql.connect("clientste.db")

    with con:
        cur = con.cursor()
        cur.execute("CREATE TABLE IF NOT EXISTS `roomstest` (`id` STRING, `number` STRING, `status` STRING)")
        
        cur.execute(f"INSERT INTO `roomstest` VALUES ('{id}', '{number}', '{status}')")
            

        con.commit()
        cur.close()

def db_emergencies(id, name, floor, note, level):

    con = sql.connect("clientste.db")
    with con:
        cur = con.cursor()
        cur.execute("CREATE TABLE IF NOT EXISTS `emergencies` (`id` STRING, `name` STRING, `floor` STRING, 'note' STRING, 'level' STRING)")
        
        cur.execute(f"INSERT INTO `emergencies` VALUES ('{id}', '{name}', '{floor}','{note}','{level}')")
            

        con.commit()
        cur.close()

def db_workerstest(id, name, surname, percode, salary):

    con = sql.connect("clientste.db")

    with con:
        cur = con.cursor()
        cur.execute("CREATE TABLE IF NOT EXISTS `workerstest` (`id` STRING, `name` STRING, `surname` STRING, `percode` STRING, `salary` STRING)")
        
        cur.execute(f"INSERT INTO `workerstest` VALUES ('{id}', '{name}', '{surname}', '{percode}', '{salary}')")
            

        con.commit()
        cur.close()

def db_clientstest(id, name, surname, percode, guest, breakfast, pool):

    con = sql.connect("clientste.db")

    with con:
        cur = con.cursor()
        cur.execute("CREATE TABLE IF NOT EXISTS `clientstest` (`id` STRING, `name` STRING, `surname` STRING, `percode` STRING, `guest` STRING, `breakfast` STRING, `pool` STRING)")
        
        cur.execute(f"INSERT INTO `clientstest` VALUES ('{id}', '{name}', '{surname}', '{percode}', '{guest}', '{breakfast}', '{pool}')")
            

        con.commit()
        cur.close()

def clear_table():

    con = sql.connect("clientste.db")
    with con:
        cur = con.cursor()
        cur.execute("TRUNCATE TABLE `clientstest`")
        
        con.commit()

def worker_window():
    global name_var, surname_var, percode_var, guestroom_var, breakfast_var, pool_var, root_main
    root_main = Toplevel()
    root_main.title("Hotel Registration stand")
    root_main.geometry("1250x820")
    root_main.resizable(False,False)
    name_var = StringVar()
    surname_var = StringVar()
    percode_var = StringVar()
    guestroom_var = StringVar()
    breakfast_var = StringVar()
    pool_var = StringVar()
    

    floor1 = Label(root_main, text="First Floor", bg='LightGrey' ,height=20, width=50, borderwidth=1, relief="solid").place(x=50, y=30)
    floor2 = Label(root_main, text="Second Floor", bg='LightGrey' ,height=20, width=50, borderwidth=1, relief="solid").place(x=450, y=30)
    floor2 = Label(root_main, text="Third Floor", bg='LightGrey' ,height=20, width=50, borderwidth=1, relief="solid").place(x=850, y=30)

    floor1_image = ImageTk.PhotoImage(Image.open("D:\Grigorijs Perecs\Logo.png"))
    floor1_image_lab = Label(root_main, image=floor1_image).place(x=30,y=30)

    Frame1 = Label(root_main, height=20, width=45, borderwidth=1, relief="solid").place(x=50, y=400)
    Frame2 = Label(root_main, height=25, width=50, borderwidth=1, relief="solid").place(x=450, y=400)

    #Buttons 1 frame


    class_3_label = Label(root_main, text = "Заселено номеров 3 класса - 0/24", font=('Arial', 12)).place(x=70, y=420)
    class_2_label = Label(root_main, text = "Заселено номеров 2 класса - 0/11", font=('Arial', 12)).place(x=70, y=460)
    class_1_label = Label(root_main, text = "Заселено номеров 1 класса - 0/1", font=('Arial', 12)).place(x=70, y=500)
    guests_label = Label(root_main, text = "Общее количество посетителей - 0", font=('Arial', 13)).place(x=70, y=550)
    Button(root_main, text = "CПИСОК", font=('Lucida Sans Unicode', 13),borderwidth=0.5, relief="solid", command=lambda:guests_list()).place(x=260, y=580)
    workers_label = Label(root_main, text = "Общее количество рабочих - 0", font=('Arial', 13)).place(x=70, y=630)
    Button(root_main, text = "CПИСОК", font=('Lucida Sans Unicode', 13),borderwidth=0.5, relief="solid", command=lambda:workers_list()).place(x=260, y=660)


    #Buttons 2 frame


    name_label = Label(root_main, text = "Имя - Григорий", font=('Arial', 13)).place(x=480, y=420)
    surname_label = Label(root_main, text = "Фамилия - Перец ", font=('Arial', 13)).place(x=480, y=460)
    percode_label = Label(root_main, text = "Персональный код - 090796-10378", font=('Arial', 13)).place(x=480, y=500)
    guestroom_label = Label(root_main, text = "Количество постоялицев - 3", font=('Arial', 13)).place(x=480, y=550)
    breakfastoption_label = Label(root_main, text = "Завтрак - да", font=('Arial', 13)).place(x=480, y=580)
    pooloption_label = Label(root_main, text = "Басейн - да", font=('Arial', 13)).place(x=480, y=630)
    numberselect_label = Label(root_main, height=3, width=30, borderwidth=1, relief="solid").place(x=520, y=675)
    numberselect_button = Button(root_main, text = "ЗАРЕГИСТРИРОВАТЬ", font=('Franklin Gothic Demi', 15),borderwidth=0, relief="solid", command=lambda: register_window()).place(x=525, y=680)


    #Emergency buttons frame

    ithelp_button = Button(root_main, text='Вызвать техника', font=('Arial', 16),borderwidth=0.5, relief="solid", bg='RoyalBlue').place(x=850, y=400)
    emergency_button = Button(root_main, text='Вызвать 112', font=('Arial', 16),borderwidth=0.5, relief="solid", bg='firebrick2').place(x=850, y=450)
    supervisor_button = Button(root_main, text='Вызвать SuperVisor', font=('Arial', 16),borderwidth=0.5, relief="solid", bg='gold2').place(x=850, y=500)

    login_button_root = Button(root_main, text="Sign in",height=2, width=10,borderwidth=0.5, relief="solid",font=("Arial", 10), command= lambda: login_window()).place(x=1115, y= 400)
    news_button_root = Button(root_main, text="News",height=2, width=10,borderwidth=0.5, relief="solid",font=("Arial", 10), command= lambda: home_page()).place(x=1115, y= 460)
    incidents_button_root = Button(root_main, text="Incidents",height=2, width=10,borderwidth=0.5, relief="solid",font=("Arial", 10), command= lambda: incidents()).place(x=1115, y= 520)


def db_login(login, password):
    
    # Connect to the database
    con = sql.connect('users.db')
    cur = con.cursor()

    # Create the logins table if it doesn't exist and insert the admin user
    cur.execute("CREATE TABLE IF NOT EXISTS `logins` (`login` STRING, `password` STRING)")
    cur.execute("INSERT OR IGNORE INTO `logins` (`login`, `password`) VALUES ('admin', 'admin')")

    # Check if the provided login and password combination exists in the database
    cur.execute(f"SELECT login, password FROM `logins` WHERE login = '{login}' AND password = '{password}'")
    result = cur.fetchone()
    print(result)
    # Print "Yes" and all records in the logins table if the login and password are valid,
    # otherwise print "Welcome"
    if result:
        
        worker_window()
        print("Yes")
        for i in con.execute('SELECT * FROM `logins`'):
            print(i)
        
    else:
        print('Welcome')

    # Close the cursor and the database connection
    cur.close()
    con.close()


available_rooms = np.arange(0,37)


def id_counter():
    id_counter_func = 0
    
    def id_counter_counter():
        nonlocal id_counter_func
        id_counter_func += 1
        return id_counter_func
    
    return id_counter_counter

id_counter_in = id_counter()
def register_guest():
    con = sql.connect('clientste.db')
    cur = con.cursor()

    id_counter_in = id_counter()
    name = name_var.get()
    surname = surname_var.get()
    percode = percode_var.get()
    guestroom = guestroom_var.get()
    breakfast = breakfast_var.get()
    pool = pool_var.get()

    cur.execute("INSERT INTO clientstest VALUES (?,?,?,?,?,?,?)", (id_counter_in(), name, surname, percode, guestroom, breakfast, pool))
    con.commit()
    con.close()

def register_window():
    # Create a new window
    hn_win = Toplevel()
    hn_win.title("Hotel Registration - Guest Room")
    hn_win.geometry("450x400")

    # Create a list of available rooms for the user to select from
    selected_room = StringVar()
    room_list = ttk.Combobox(hn_win, textvariable=selected_room, cursor='arrow', height=10, width=15)
    room_list['values'] = [room for room in available_rooms if room not in ['{', '}']]
    room_list['state'] = 'readonly'
    room_list.place(x=210, y=70)

    # Create labels and input fields for the user to enter their name, surname, personal code, guest room, breakfast preference, and pool preference
    Label(hn_win, text="Name:").place(x=100, y=100)
    guest_name_input = Entry(hn_win, textvariable = name_var, width=20, bd = 5, borderwidth=1, relief="solid")
    guest_name_input.place(x = 210, y = 100)

    Label(hn_win, text="Surname:").place(x=100, y=130)
    guest_surname_input = Entry(hn_win, textvariable = surname_var, width=20, bd = 5, borderwidth=1, relief="solid")
    guest_surname_input.place(x = 210, y = 130)

    Label(hn_win, text="Personal Code:").place(x=100, y=160)
    guest_percode_input = Entry(hn_win, textvariable = percode_var, width=20, bd = 5, borderwidth=1, relief="solid")
    guest_percode_input.place(x = 210, y = 160)

    Label(hn_win, text="Guest Room:").place(x=100, y=190)
    guest_guestroom_guest_input = Entry(hn_win, textvariable = guestroom_var, width=20, bd = 5, borderwidth=1, relief="solid")
    guest_guestroom_guest_input.place(x = 210, y = 190)

    Label(hn_win, text="Breakfast Prefer:").place(x=100, y=210)
    guest_breakfast_input = Entry(hn_win, textvariable = breakfast_var, width=20, bd = 5, borderwidth=1, relief="solid")
    guest_breakfast_input.place(x = 210, y = 210)

    Label(hn_win, text="Pool Prefer:").place(x=100, y=240)
    guest_pool_input = Entry(hn_win, textvariable = pool_var, width=20, bd = 5, borderwidth=1, relief="solid")
    guest_pool_input.place(x = 210, y = 240)

    # Create a button for the user to register their information
    register_button = Button(hn_win, text = "REGISTER", font=('Franklin Gothic Demi', 15),borderwidth=0.5, relief="solid", command=lambda: register_guest())
    register_button.place(x=120, y=270)

    # Run the window's main loop
    hn_win.mainloop()

def incidents():
    incidents_win = Toplevel()
    incidents_win.title("Incidents")
    incidents_win.geometry("650x500")
    con = sql.connect('incidents.db')
    cur = con.cursor()


    cur.execute("CREATE TABLE IF NOT EXISTS Incidents (id INTEGER PRIMARY KEY, name TEXT, floor INTEGER, note TEXT, level INTEGER, status TEXT)")


    dataforfun = cur.execute("SELECT * FROM `Incidents` ORDER BY level ASC")
    rows = cur.fetchall()
    table = ttk.Treeview(incidents_win, columns=("id", "name", "floor", "note", "level", "status"), show="headings")


    table.heading("#0", text="")
    table.heading("id", text="ID")
    table.heading("name", text="Name")
    table.heading("floor", text="Floor")
    table.heading("note", text="Note")
    table.heading("level", text="Level")
    table.heading("status", text="Status")


    for i, row in enumerate(rows):
        table.insert("", "end", values=row)


    table.column("#0", minwidth=0, width=0, stretch=NO)
    table.column("id", minwidth=0, width=50, stretch=NO)
    table.column("name", minwidth=0, width=150, stretch=NO)
    table.column("floor", minwidth=0, width=50, stretch=NO)
    table.column("note", minwidth=0, width=200, stretch=NO)
    table.column("level", minwidth=0, width=50, stretch=NO)
    table.column("status", minwidth=0, width=100, stretch=NO)
    table.pack(side=LEFT, fill=BOTH)


    name_entry = Entry(incidents_win, width=20)
    name_entry.place(x=10, y=450)
    floor_entry = Entry(incidents_win, width=20)
    floor_entry.place(x=110, y=450)
    note_entry = Entry(incidents_win, width=20)
    note_entry.place(x=210, y=450)
    level_entry = Entry(incidents_win, width=20)
    level_entry.place(x=310, y=450)
    status_entry = Entry(incidents_win, width=20)
    status_entry.place(x=410, y=450)


    submit_button = Button(incidents_win, text="Submit", command=lambda: insert_incident(name_entry.get(), floor_entry.get(), note_entry.get(), level_entry.get(), status_entry.get()))
    submit_button.place(x=520, y=450)
    id_combobox = Combobox(incidents_win, width=20)
    id_combobox.place(x=10, y=400)

    cur.execute("SELECT id FROM Incidents")
    incident_ids = cur.fetchall()


    id_combobox['values'] = [id[0] for id in incident_ids]

    delete_button = Button(incidents_win, text="Delete", command=lambda: delete_incident(id_combobox.get()))
    delete_button.place(x=200, y=400)
#####################################################################################################################################
    id_entry = Entry(incidents_win, width=20)
    id_entry.place(x=10, y=350)
    name_entry = Entry(incidents_win, width=20)
    name_entry.place(x=110, y=350)
    note_entry = Entry(incidents_win, width=20)
    note_entry.place(x=210, y=350)


    edit_button = Button(incidents_win, text="Edit", command=lambda: update_incident(id_entry.get(), name_entry.get(), note_entry.get()))
    edit_button.place(x=310, y=350)

    def update_incident(incident_id, name, note):

        cur.execute("UPDATE Incidents SET id=?, name=?, note=? WHERE id=?", (incident_id, name, note, incident_id))
        con.commit()
        

    def delete_incident(incident_id):

        cur.execute("DELETE FROM Incidents WHERE id=?", (incident_id,))
        con.commit()

    def insert_incident(name, floor, note, level, status):

        cur.execute("INSERT INTO Incidents (name, floor, note, level, status) VALUES (?, ?, ?, ?, ?)", (name, floor, note, level, status))
        con.commit()

def login_user(login, password):
    db_login(login, password)

def login_window():
    login_var = StringVar()
    password_var = StringVar()
    login_win = Toplevel()
    login_win.title("Sign in")
    login_win.geometry("350x200")
    admin_button = Button(login_win, text = "Admin", font=('Franklin Gothic Demi', 15),borderwidth=0, relief="solid", command=lambda: register_window())
    admin_button.place(x=0, y=0)
    Login_label = Label(login_win, text = "Login:  ", font=('Arial', 11))
    Login_label.place(x=40, y=10)
    Login_entry = Entry(login_win, textvariable=login_var, width=15, bg='LightGrey')
    Login_entry.place(x= 120, y = 10)
    Password_label = Label(login_win, text = "Password:  ", font=('Arial', 11))
    Password_label.place(x=40, y=60)
    Password_entry = Entry(login_win, textvariable=password_var, width=15, bg='LightGrey')
    Password_entry.place(x= 120, y = 60)
    Enter_button = Button(login_win, text = "ENTER",height = 1, width= 8, font=('Franklin Gothic Demi', 15),borderwidth=1, relief="solid", command=lambda: login_user(login_var.get(), password_var.get()))
    Enter_button.place(x = 40, y = 100)

def update_time():
    # Get the current time
    current_time = time.strftime("%H:%M:%S")
    # Update the time label
    time_label.config(text=current_time)
    # Schedule the next update
    gwin_win.after(1000, update_time)

def window_main():
    global time_label, gwin_win
    gwin_win = Tk()
    gwin_win.title("Hotel ALTAIR")
    gwin_win.geometry("1200x800")
    con = sql.connect('clientste.db')
    cur = con.cursor()

    login_btn_work = Button(gwin_win, text="Sign in",height=2, width=10,borderwidth=0.5, relief="solid",font=("Arial", 10), command= lambda: login_window()).place(x=1090, y=25)
    time_label = Label(gwin_win, text="00:00:00", font=('Arial', 18))
    time_label.place(x=20, y=20)
    img1 = ImageTk.PhotoImage(Image.open("D:\Grigorijs Perecs\Logo.png").resize((400,200)))
    myLabel = Label(gwin_win, image = img1).place(x=370, y=20)


    update_time()
    gwin_win.mainloop()

def guests_list():
    glist_win = Toplevel()
    glist_win.title("Guests list")
    glist_win.geometry("400x800")
    con = sql.connect('clientste.db')
    cur = con.cursor()

    

    cur.execute("SELECT * FROM `clientstest`")
    rows = cur.fetchall()
    for i, row in enumerate(rows):
        label_text = ", ".join([str(x) for x in row])
        Label(glist_win, text=label_text, font=("Arial", 13)).place(x=20, y=5 + 30 * i)
        
def workers_list():
    wlist_win = Toplevel()
    wlist_win.title("Workers list")
    wlist_win.geometry("400x800")
    con = sql.connect('clientste.db')
    cur = con.cursor()

    dataforfun = cur.execute("SELECT * FROM `workerstest`")
    rows = cur.fetchall()
    for i in range(len(rows)):
        Label(wlist_win, text=rows[i], font=("Arial", 13)).place(x = 20, y = 5+30*i)

def home_page():
    # Create a new tab for the Home page
    home_tab = Toplevel()
    home_tab.title("News")
    home_tab.geometry("400x800")

    # Connect to the database
    con = sql.connect('news.db')
    cur = con.cursor()
    cur.execute("SELECT title, text, date FROM `news` ORDER BY date")
    rows = cur.fetchall()
    for i in range(len(rows)):
        label_text = str(rows[i]).replace("{", "").replace("}", "")
        Label(home_tab, text=label_text, font=("Arial", 13)).place(x = 20, y = 5+30*i)
    cur.close()
    con.close()


#backround place

window_main()



