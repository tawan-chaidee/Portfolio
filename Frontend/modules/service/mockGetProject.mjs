import {Project} from '../model/project.mjs';



  
const mockData = [
  {
    "name": "Project A",
    "description": "Description for Project A",
    "startDate": "2024-01-01",
    "endDate": "2024-01-31",
    "status": "In Progress",
    "duration": "30 days"
  },
  {
    "name": "Project B",
    "description": "Description for Project B",
    "startDate": "2024-02-01",
    "endDate": "2024-02-28",
    "status": "Completed",
    "duration": "27 days"
  },
  {
    "name": "Project C",
    "description": "Description for Project C",
    "startDate": "2024-03-01",
    "endDate": "2024-03-31",
    "status": "Planned",
    "duration": "30 days"
  }
];
  
export function getProject() {
  const projects = mockData.map(data => new Project(
    data.name,
    data.description,
    data.startDate,
    data.endDate,
    data.status
  ));
  return projects
}

