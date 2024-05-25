


import { getProject } from "./modules/service/mockGetProject.mjs";

function handleClick() {
    const projects = getProject()

    projects.map(project => {
        project.displayDetails()
    })
}

document.getElementById('myButton').addEventListener('click', handleClick);
  