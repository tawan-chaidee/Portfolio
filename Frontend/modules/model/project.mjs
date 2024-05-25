export class Project {
  constructor(name, description, startDate, endDate, status) {
    this.name = name;
    this.description = description;
    this.startDate = startDate;
    this.endDate = endDate;
    this.status = status;
  }


  displayDetails() {
    console.log(`Project Name: ${this.name}`);
    console.log(`Description: ${this.description}`);
    console.log(`Start Date: ${this.startDate}`);
    console.log(`End Date: ${this.endDate}`);
    console.log(`Status: ${this.status}`);
    // console.log(`Duration: ${this.getDuration()} days`);
  }




}
