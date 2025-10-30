# playcanvas-script-share

A command-line utility to quickly push and share identical scripts across multiple PlayCanvas projects.

## 💬 Overview

Managing the same script file across dozens of projects is tedious. This tool automates the process by running the official PlayCanvas `pcsync` utility for every project you define, ensuring all your projects receive the latest version of your shared scripts from a single local folder.

This tool works by leveraging `pcsync`'s ability to merge a global config file (for your API key) with multiple local config files (for each project's ID).

## ⚙️ How It Works

1.  **Global Config:** `pcsync` first reads the `.pcconfig` file in your system's User directory (e.g., `C:\Users\yourName`) to get your master `PLAYCANVAS_API_KEY` and the `PLAYCANVAS_TARGET_DIR` (the path to your `pcSync` folder).
2.  **Local Configs:** This tool (`sync.exe`) reads all project files (e.g., `project_A.json`) from the `Configs` folder.
3.  **Execution:** For each project, the tool:
    a.  Copies `Configs/project_A.json` to the `pcSync` folder and renames it to `pcconfig.json`. This file *only* contains the Project ID and Branch ID.
    b.  Runs `pcsync pushAll --yes` from within the `pcSync` directory.
    c.  `pcsync` sees both the **global `.pcconfig` (with API key)** and the **local `pcconfig.json` (with Project ID)**, merges them, and pushes the scripts.
    d.  The tool repeats this for `project_B.json`, `project_C.json`, and so on.

## 📥 Prerequisites

  * The **`sync.exe`** executable from this repository.
  * The official PlayCanvas **`pcsync`** command-line tool.
  * A **PlayCanvas API Key** (Personal Access Token).

## 🚀 Setup Instructions

### Step 1: Create Your Folder Structure

Create a main folder and organize it exactly like this:

```
playcanvas-sync-tool/
│
├── Configs/
│   ├── project_one.json
│   ├── project_two.json
│   └── (any_other_project.json...)
│
├── pcSync/
│   ├── your_shared_script.js
│   ├── your_shared_script_folder
│   ├── pcconfig.json
│   └── pcignore.txt
│
└── sync.exe
```

  * **`Configs/`**: Holds the *project-specific* configs (ID and Branch).
  * **`pcSync/`**: This is the "source" folder. Place all the scripts or folders you want to upload here.
  * **`sync.exe`**: The executable from this repository.

### Step 2: Create Global Config (`.pcconfig`)

This file holds your secret API key and tells `pcsync` where your `pcSync` folder is. You only need to do this **one time**.

1.  Navigate to your User's home directory (e.g., `C:\Users\yourName` on Windows or `~` on Mac/Linux).

2.  Create a new file named exactly **`.pcconfig`** (with no file extension).

3.  Open it and paste the following, modifying the values:

    ```json
    {
      "PLAYCANVAS_API_KEY": "your-api-key-goes-here",
      "PLAYCANVAS_TARGET_DIR": "C:\\path\\to\\your\\playcanvas-script-share\\pcSync",
      "PLAYCANVAS_BAD_FILE_REG": "^\\.|~$",
      "PLAYCANVAS_BAD_FOLDER_REG": "\\."
    }
    ```

    > **Warning:** `PLAYCANVAS_TARGET_DIR` **must be the full, absolute path** to your `pcSync` folder.

    >   * Windows Example: `C:\\Users\\MyUser\\Desktop\\playcanvas-script-share\\pcSync` (Note the double backslashes `\\` for JSON)
    >   * Mac/Linux Example: `/Users/MyUser/Desktop/playcanvas-script-share/pcSync`

### Step 3: Create Project Configs (in `Configs/`)

Next, you must create one `.json` file inside the `Configs` folder for *every* project you want to sync.

1.  Open your PlayCanvas project in the Editor.

2.  Open your browser's **Developer Tools** (F12) and go to the **Console** tab.

3.  Paste and run the following command to copy your project IDs:

    ```javascript
    copy({
      PLAYCANVAS_BRANCH_ID: config.self.branch.id,
      PLAYCANVAS_PROJECT_ID: config.project.id
    })
    ```

4.  In your `Configs` folder, create a new file (e.g., `my_first_project.json`).

5.  Paste the content from your clipboard. The file should look like this:

    ```json
    {
      "PLAYCANVAS_BRANCH_ID": "your-branch-id-from-console",
      "PLAYCANVAS_PROJECT_ID": 123456
    }
    ```

6.  Repeat this for all your projects, creating a new `.json` file for each one in the `Configs` folder.

## ▶️ How to Use

Once your folder structure and all config files are in place:

1.  Add or update the scripts you want to share in the `pcSync/` folder.
2.  Double-click and run **`sync.exe`**.

The tool will run and push your scripts to every project defined in your `Configs` folder.
