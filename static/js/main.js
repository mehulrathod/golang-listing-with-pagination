let currentPage = 1;
const getList = (page = 1, pageSize  = 10) => {
    fetch("http://localhost:8080/api/v1/employee-list", {method: "POST", body: JSON.stringify({page, pageSize})})
        .then(response => response.json())
        .then(prepareList);
}

const prepareList = (result) => {
    preparePaging();
    const listContainer = document.querySelector(".employee-list");
    listContainer.innerHTML = '<tr>' + result.data.map(employee => `<td>${employee.id}</td><td>${employee.name}</td><td>${employee.city}</td>`).join("</tr><tr>")
}
const pageChange = (pageNum) => {
    currentPage = pageNum;
    getList(pageNum);
};
const preparePaging = () => {
    const pagingContainer = document.querySelector(".pagination");
    // const pageSize = 10;
    // const numOfPages = Math.floor(1000/pageSize);

    let html = '';
    if(currentPage >= 3) {
        html += `<li class="page-item ">
                    <a class="page-link" href="javascript:void(0)" onclick="pageChange(${currentPage - 1})">Previous</a>
                </li>`;
    }
    for(let i=currentPage === 1 ? 1 : currentPage - 1; i<= currentPage + (currentPage === 1 ? 2 : 1); i++) {
        html += `<li class="page-item ${currentPage === i ? 'active' : ''}">
                    <a class="page-link" href="javascript:void(0)" onclick="pageChange(${i})">${i}</a>
                </li>`;
    }
    if(3 <= currentPage ) {
        html += `<li class="page-item ">
                    <a class="page-link" href="javascript:void(0)" onclick="pageChange(${currentPage + 1})">Next</a>
                </li>`;
    }

    pagingContainer.innerHTML = html;
}

getList();