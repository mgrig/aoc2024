##part 1
mat = load('../resources/01.txt');
mat = sort(mat);
sum(abs(mat(:, 1) - mat(:, 2)))

##part 2
[uniques, ~, idx] = unique(mat(:,2));
counts = accumarray(idx(:), 1);
h = [uniques, counts];
elementMap = containers.Map(h(:,1), h(:,2));

counts_for_list1 = zeros(size(mat,1), 1);
for i = 1:length(mat(:,1))
  x = mat(i);
  if isKey(elementMap, x)
    counts_for_list1(i) = elementMap(x);
  endif
end

dot(mat(:,1), counts_for_list1)
