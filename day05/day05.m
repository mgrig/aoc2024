%{
See the rules as an edge in a directed graph.
So 'A|B' means "page A must come before page B" and translates to a A -> B directed edge.

Use the Floyd-Warshall algorithm to calculate the shortest path lengths between
every pair of vertices in the graph.

The rest is just about evaluating the result of the computed FW matrix.

The needed helper functions were written with AI, so beware "there be dragons in there",
but they seem to do the right thing... after enough iterations.
%}
function day05
  dist = load_directed_graph('05_rules.txt', 99);
  dist(1:size(dist,1)+1:end) = 0; # set main diagonal to 0

  updates = load_comma_separated_struct_array('05_updates.txt');

  sum1 = 0;
  sum2 = 0;
  for r = 1:length(updates)
    line = updates(r).data;
    right_order = true;

    dist_copy = Inf(99, 99);
##    keep only rules among the numbers in 'line'
    for ir = 1:length(line)
      for ic = 1:length(line)
        dist_copy(line(ir), line(ic)) = dist(line(ir), line(ic));
      endfor
    endfor

    dist_copy = compute_fw(dist_copy);

    for c = 1:length(line)-1
     d = dist_copy(line(c), line(c+1));
     if isinf(d)
        right_order = false;
        break
      endif
    endfor

    if right_order
      mid_value = line((1+length(line))/2);
      sum1 = sum1 + mid_value;
    else
##      sort the line
      cmp_func = @(x, y) ( 2 * isinf(dist_copy(x, y)) - 1 );
      sorted_line = custom_sort_unique(line, cmp_func);

      mid_value = sorted_line((1+length(sorted_line))/2);
      sum2 = sum2 + mid_value;
    endif
  endfor

  fprintf("sum1 = %d\n", sum1)
  fprintf("sum2 = %d\n", sum2)
end

function dist = compute_fw(dist)
## Use Floyd-Warshall algorithm to calculate the shortest path lengths between every pair of vertices in a graph
  for k = 1:size(dist,1)
    dist = min (dist, dist(:,k) + dist(k,:));
  endfor
end

function adj_matrix = load_directed_graph(filename, n)

    % Check if the file exists
    if ~exist(filename, 'file')
        error('File not found: %s', filename);
    end

    % Open the file for reading
    fid = fopen(filename, 'r');
    if fid == -1
        error('Cannot open file: %s', filename);
    end

    % Read the entire file using textscan
    % Each line is expected to have two integers separated by '|'
    data = textscan(fid, '%d|%d', 'Delimiter', '|', 'CollectOutput', true, ...
                   'TreatAsEmpty', {'', 'Inf'}, 'CommentStyle', '#');
    fclose(fid);

    % Extract the edges from the scanned data
    edges = data{1}; % Nx2 matrix where each row is [from, to]

    % Check if any valid edges were found
    if isempty(edges)
        warning('No valid edges found in the file.');
        adj_matrix = [];
        return;
    end

    % Extract 'from' and 'to' nodes
    from_nodes = edges(:, 1);
    to_nodes = edges(:, 2);

    % Determine the maximum node number to size the adjacency matrix
##    max_from = max(from_nodes);
##    max_to = max(to_nodes);
##    max_node = max(max_from, max_to);
    max_node = n;

    % Initialize the adjacency matrix with Inf
    adj_matrix = Inf(max_node, max_node);

    % Convert 'from' and 'to' node pairs to linear indices
    % and set the corresponding entries in the adjacency matrix to 1
    indices = sub2ind(size(adj_matrix), from_nodes, to_nodes);
    adj_matrix(indices) = 1;

##    % Optional: Display a message indicating successful loading
##    fprintf('Loaded directed graph from "%s".\n', filename);
##    fprintf('Number of nodes: %d\n', max_node);
##    fprintf('Number of edges: %d\n', length(indices));
end

function matrix_struct = load_comma_separated_struct_array(filename)
    %LOAD_COMMA_SEPARATED_STRUCT_ARRAY Loads a variable-length matrix from a file.  updates = load_comma_separated_struct_array('05_updates.txt');

    %   matrix_struct = LOAD_COMMA_SEPARATED_STRUCT_ARRAY(filename) reads a file where each
    %   line contains a comma-separated list of integers with no spaces. Each line can have
    %   a different number of integers. It returns a struct array where each struct corresponds
    %   to a row in the file, containing a field 'data' with the row vector of integers.
    %
    %   Input:
    %       filename - String representing the path to the input file.
    %
    %   Output:
    %       matrix_struct - 1xM struct array where each struct has a field 'data'
    %                       containing a 1xN_i integer row vector.
    %
    %   Example:
    %       If the file "data_variable.csv" contains:
    %           1,2,3
    %           4,5
    %           6,7,8,9
    %       Then, calling:
    %           mat_struct = load_comma_separated_struct_array('data_variable.csv');
    %       Will return:
    %           mat_struct(1).data = [1 2 3]
    %           mat_struct(2).data = [4 5]
    %           mat_struct(3).data = [6 7 8 9]

    % Check if the file exists
    if ~exist(filename, 'file')
        error('File not found: %s', filename);
    end

    % Open the file for reading
    fid = fopen(filename, 'r');
    if fid == -1
        error('Cannot open file: %s', filename);
    end

    % Initialize an empty struct array
    matrix_struct = struct('data', {});

    % Read the file line by line
    line_num = 0;
    while ~feof(fid)
        line = fgetl(fid);
        line_num = line_num + 1;

        % Trim whitespace from the line
        line = strtrim(line);

        % Skip empty lines
        if isempty(line)
            continue;
        endif

        % Optionally, skip comment lines (starting with #)Floyd-Warshall algorithm to calculate the shortest path lengths between every pair of vertices in a graph
        if startsWith(line, '#')
            continue;
        endif

        % Split the line by commas
        parts = strsplit(line, ',');

        % Convert each part to an integer
        try
            numbers = cellfun(@str2num, parts, 'UniformOutput', true);
        catch
            fclose(fid);
            error('Non-integer value found on line %d: "%s".', line_num, line);
        end_try_catch

        % Check for successful conversion
        if any(isnan(numbers))
            fclose(fid);
            error('Invalid integer found on line %d: "%s".', line_num, line);
        endif

        % Append to the struct array
        matrix_struct(end + 1).data = numbers;
    endwhile

    fclose(fid);

    % Check if any data was read
    if isempty(matrix_struct)
        warning('No data found in the file: %s.', filename);
    endif
end
function sorted = custom_sort_unique(A, cmp_func)
    %CUSTOM_SORT_UNIQUE Sorts a row vector using a custom comparison function.
    %   sorted = CUSTOM_SORT_UNIQUE(A, cmp_func) takes a 1D array A with unique
    %   elements and a comparison function cmp_func, and returns a new row vector
    %   sorted based on cmp_func.
    %
    %   Inputs:
    %       A         - 1D array to be sorted. All elements must be unique.
    %       cmp_func  - Function handle that takes two inputs and returns:
    %                   - Negative if first < second
    %                   - Positive if first > second
    %
    %   Output:
    %       sorted    - Sorted row vector based on cmp_func.
    %
    %   Example:
    %       % Comparator to sort numbers by absolute value
    %       cmp_abs = @(x, y) abs(x) - abs(y);
    %       sorted_array = custom_sort_unique([-3, 1, -2, 4, 0], cmp_abs);
    %       % sorted_array = [0, 1, -2, -3, 4]

    % Ensure A is a row vector
    A = A(:).';

    % Validate that A is a vector
    if ~isvector(A)
        error('Input A must be a 1D array.');
    end

    % Validate that all elements are unique
    if length(unique(A)) != length(A)
        error('All elements in input array A must be unique.');
    end

    % Base case: empty array or single element is already sorted
    if isempty(A) || length(A) == 1
        sorted = A;
        return;
    endif

    % Choose the pivot (last element)
    pivot = A(end);

    % Initialize containers for partitions
    less = [];
    greater = [];

    % Partition the array based on the comparison with pivot
    for i = 1:length(A)-1  % Exclude the pivot itself
        result = cmp_func(A(i), pivot);
        if result < 0
            less(end + 1) = A(i); %#ok<AGROW>
        else
            greater(end + 1) = A(i); %#ok<AGROW>
        endif
    endfor

    % Recursively sort the partitions
    if isempty(less)
        sorted_less = [];
    else
        sorted_less = custom_sort_unique(less, cmp_func);
    endif

    if isempty(greater)
        sorted_greater = [];
    else
        sorted_greater = custom_sort_unique(greater, cmp_func);
    endif

    % Concatenate the results: [sorted_less, pivot, sorted_greater]
    sorted = [sorted_less, pivot, sorted_greater];
endfunction

