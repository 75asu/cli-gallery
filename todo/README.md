### A todo app with CRUD operations

**Demo**

```shell
gitpod /workspace/cli-gallery/todo (main) $ go build && go install
gitpod /workspace/cli-gallery/todo (main) $ todo add my1stTask --description "Tide my room" --deadline "30-10-2022"
Using config file: /workspace/cli-gallery/todo/.todos.yaml
One todo added
gitpod /workspace/cli-gallery/todo (main) $ todo add my2ndTask --description "Read xyz book" --deadline "31-10-2022"
Using config file: /workspace/cli-gallery/todo/.todos.yaml
One todo added
gitpod /workspace/cli-gallery/todo (main) $ todo list
Using config file: /workspace/cli-gallery/todo/.todos.yaml
[
 {
  "Name": "my1stTask",
  "Description": "This is an updated description",
  "Deadline": "30-10-2022"
 },
 {
  "Name": "my1stTask",
  "Description": "Tide my room",
  "Deadline": "30-10-2022"
 },
 {
  "Name": "my2ndTask",
  "Description": "Read xyz book",
  "Deadline": "31-10-2022"
 }
]
gitpod /workspace/cli-gallery/todo (main) $ todo remove my2ndTask
Using config file: /workspace/cli-gallery/todo/.todos.yaml
todo removed
gitpod /workspace/cli-gallery/todo (main) $ todo list
Using config file: /workspace/cli-gallery/todo/.todos.yaml
[
 {
  "Name": "my1stTask",
  "Description": "This is an updated description",
  "Deadline": "30-10-2022"
 },
 {
  "Name": "my1stTask",
  "Description": "Tide my room",
  "Deadline": "30-10-2022"
 }
]
gitpod /workspace/cli-gallery/todo (main) $ todo update my1stTask --description "This is an updated description"
Using config file: /workspace/cli-gallery/todo/.todos.yaml
todo updated
gitpod /workspace/cli-gallery/todo (main) $ todo list
Using config file: /workspace/cli-gallery/todo/.todos.yaml
[
 {
  "Name": "my1stTask",
  "Description": "This is an updated description",
  "Deadline": "30-10-2022"
 },
 {
  "Name": "my1stTask",
  "Description": "This is an updated description",
  "Deadline": "30-10-2022"
 }
]
gitpod /workspace/cli-gallery/todo (main) $ 
```