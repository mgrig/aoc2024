function day02()
  files = {'../resources/02_count_4.txt', '../resources/02_count_5.txt', '../resources/02_count_6.txt', '../resources/02_count_7.txt'};
  count = 0;
  for i = 1:length(files)
    mat = dlmread(files{i}, ' ');
    count = count + countSafe(mat);
  end
  count
end

function c = countSafe(mat)
  d = diff(mat, 1, 2);
  a = abs(d);
  c = sum((all(d > 0, 2) | all(d < 0, 2)) & all(a > 0, 2) & all(a <= 3, 2));
end
