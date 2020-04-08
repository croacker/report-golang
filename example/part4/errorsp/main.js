const employee = {
    name: 'name',
    salary: 100
}

const checkSalary = function (v) {
    if(v > 10) {
        throw 'too much'
    } else {
        return 'good'
    }
}

try{
    const result = checkSalary(employee.salary)
    console.log(result)
} catch (e) {
    console.log(e)
}