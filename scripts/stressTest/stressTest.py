from locust import User, task, between, events, TaskSet, HttpUser
from locust.contrib.fasthttp import FastHttpUser
from users import users_info
import random
import time

class storage(object):
    token = ''

class UserBehaviour(FastHttpUser):
    # wait_time = between(0.5, 5)

    def get_user(self):
        user = random.choice(users_info)
        return user

    def on_start(self):
        self.user = self.get_user()
        response =  self.client.post("/login", json = {
            "username": self.user[1],
            "password": '245'}) 
        self.token = response.json().get("access_token")
        
    @task
    def my_task(self):
        self.client.get("/")

    @task
    def my_profile(self):
        self.client.get("/profile", headers= {'Content-Type': 'application/json', 'Authorization': "Bearer " + self.token})

    @task
    def show_picture(self):
        self.client.get("/showProfilePicture", headers= {'Content-Type': 'application/json', 'Authorization': "Bearer " + self.token})

    @task
    def profile_update(self):
        self.client.post("/profile/update", headers= {'Content-Type': 'application/json', 'Authorization': "Bearer " + self.token}, 
        json = {"name": self.user[0]})
    
    @task
    def change_password(self):
        self.client.post("/changePassword", headers= {'Content-Type': 'application/json', 'Authorization': "Bearer " + self.token}, 
        json = {"password": "245"})

    # @task 
    # def upload(self):
    #     attach = open('picture.png', 'r')
    #     self.client.post("/uploadProfilePicture", headers= {'Authorization': "Bearer " + self.token}, 
    #     files = {"myFile": attach})

    def on_stop(self):
        self.client.get("/logout", headers= {'Authorization': "Bearer " + self.token})