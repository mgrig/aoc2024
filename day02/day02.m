function day02()
  files = {'../resources/02_count_4.txt', '../resources/02_count_5.txt', '../resources/02_count_6.txt', '../resources/02_count_7.txt'};
##  files = {'../resources/02_test.txt'};
  part1 = 0;
  part2 = 0;
  for i = 1:length(files)
    mat = dlmread(files{i}, ' ');
    part1 = part1 + countSafe(mat);
    part2 = part2 + countSafeLenient(mat);
  end
  part1
  part2
end

function [c, mask] = countSafe(mat)
  d = diff(mat, 1, 2);
  a = abs(d);
  mask = (all(d > 0, 2) | all(d < 0, 2)) & all(a > 0, 2) & all(a <= 3, 2);
  c = sum(mask);
end

function cLenient = countSafeLenient(mat)
  [cLenient, mask] = countSafe(mat);

  mat = mat(mask == 0, :); % keep only unsafe rows

  % Remove each column individually, count safe rows, and continue ONLY with the remaining unsafe rows
  for i = 1:size(mat,2)
    mat_small = mat;
    mat_small(:, i) = []; % remove column i from a copy of mat

    [c, mask] = countSafe(mat_small);
    cLenient = cLenient + c;

    mat = mat(mask == 0, :); % keep only remaining unsafe rows
  end
end
