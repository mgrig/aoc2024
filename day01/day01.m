##part 1
mat = load('../resources/01.txt');
mat = sort(mat);
sum(abs(diff(mat, 1, 2)))

##part 2
[uniques, ~, idx] = unique(mat(:,2));
counts = accumarray(idx(:), 1);
elementMap = containers.Map(uniques, counts);

counts_for_list1 = zeros(size(mat,1), 1);
for i = 1:length(mat(:,1))
  x = mat(i);
  if isKey(elementMap, x)
    counts_for_list1(i) = elementMap(x);
  endif
end

format long g
dot(mat(:,1), counts_for_list1)
